package inject

import (
    "database/sql"
    "fmt"
    "github.com/BurntSushi/toml"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    s3Session "github.com/aws/aws-sdk-go/aws/session"
    "github.com/go-playground/validator/v10"
    "github.com/gorilla/sessions"
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/cfg"
    "github.com/sia-fl/medicine/src/proto"
    "github.com/sia-fl/medicine/src/util"
    "go.uber.org/dig"
    "google.golang.org/grpc"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
    "regexp"
    "sync"
)

type (
    CfgPath struct {
        Path string
    }
    AppAdmin struct {
        dig.In
        App *iris.Application `name:"AppAdmin"`
    }
)

var (
    initLock  = &sync.Once{}
    container *dig.Container
)

func Inject() *dig.Container {
    initLock.Do(func() {
        container = dig.New()
        util.CE(container.Provide(genCfg))
        util.CE(container.Provide(genDb))
        util.CE(container.Provide(genOrm))
        util.CE(container.Provide(genIdsRpcSer))
        util.CE(container.Provide(genSessions))
        util.CE(container.Provide(genValidator))
        util.CE(container.Provide(genS3Session))
        genAppAdmin(container)
    })
    return container
}

func genCfg(c *CfgPath) (*cfg.App, *cfg.Database, *cfg.Service) {
    var path string
    if c == nil {
        path = "./cfg/used.toml"
    } else {
        path = c.Path
    }
    var cf cfg.App
    util.CE(toml.DecodeFile(path, &cf))
    return &cf, cf.Database, cf.Service
}

func genDb(d *cfg.Database) *sql.DB {
    db, err := sql.Open(d.Drive, d.Conn)
    util.CE(err)
    return db
}

func genOrm(db *sql.DB, d *cfg.Database) *gorm.DB {
    orm, err := gorm.Open(
        mysql.New(mysql.Config{Conn: db}),
        &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: d.Prefix}},
    )
    util.CE(err)
    return orm
}

func genIdsRpcSer(s *cfg.Service) (proto.IdInMapServicesClient, proto.IdServicesClient) {
    conn, err := grpc.Dial(fmt.Sprintf("%s:%s", s.Ids.Host, s.Ids.Port), grpc.WithInsecure())
    util.CE(err)
    return proto.NewIdInMapServicesClient(conn), proto.NewIdServicesClient(conn)
}

func genSessions(cf *cfg.App) *sessions.CookieStore {
    return sessions.NewCookieStore([]byte(cf.Secret))
}

func genValidator() *validator.Validate {
    validate := validator.New()
    _ = validate.RegisterValidation("md5", func(f validator.FieldLevel) bool {
        match, _ := regexp.MatchString("^[a-z0-9]{32}$", f.Field().String())
        return match
    })
    _ = validate.RegisterValidation("md5filename", func(f validator.FieldLevel) bool {
        match, _ := regexp.MatchString("^[a-z0-9]{32}.(^jpg|jpeg|png)$", f.Field().String())
        return match
    })
    _ = validate.RegisterValidation("phone", func(f validator.FieldLevel) bool {
        match, _ := regexp.MatchString("^1[3-9]\\d{9}$", f.Field().String())
        return match
    })
    _ = validate.RegisterValidation("zh", func(f validator.FieldLevel) bool {
        match, _ := regexp.MatchString("^[\\u4e00-\\u9fa5]+$", f.Field().String())
        return match
    })
    _ = validate.RegisterValidation("en", func(f validator.FieldLevel) bool {
        match, _ := regexp.MatchString("^[a-zA-Z\\S]+$", f.Field().String())
        return match
    })
    return validate
}

func genS3Session(c *cfg.App) *s3Session.Session {
    sess, err := s3Session.NewSession(&aws.Config{
        Region:      aws.String(c.S3.Region),
        Credentials: credentials.NewStaticCredentials(c.S3.S3ID, c.S3.S3Secret, ""),
        DisableSSL:  aws.Bool(c.S3.DisableSSL),
        Endpoint:    aws.String(c.S3.Schema + c.S3.Endpoint),
    })
    util.CE(err)
    _, err = sess.Config.Credentials.Get()
    util.CE(err)
    return sess
}
