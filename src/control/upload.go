package control

import (
    "errors"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    s3Session "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/util"
    "net/http"
    "regexp"
    "time"
)

type UploadProps struct {
    Type   string
    Path   string
    Bucket string
    S3Sess *s3Session.Session
}

type uploadForm struct {
    FileMd5 string `json:"file_md_5"`
    FileExt string `json:"file_ext"`
}

var extErr = errors.New("错误的文件格式")

func CheckExt(t string, ext string) {
    switch t {
    case "image":
        match, _ := regexp.MatchString("^(jpeg|jpg|png)$", ext)
        if !match {
            panic(extErr)
        }
        break
    default:
        panic(extErr)
    }
}

func Upload(ctx iris.Context, ps *UploadProps) {
    var form uploadForm
    util.CE(ctx.ReadJSON(&form))
    CheckExt(ps.Type, form.FileExt)
    EId := EId(ctx)
    key := fmt.Sprintf("%s/%d/%s.%s", ps.Path, EId, form.FileMd5, form.FileExt)
    svc := s3.New(ps.S3Sess)
    req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
        Bucket: aws.String(ps.Bucket),
        Key:    aws.String(key),
    })
    url, err := req.Presign(30 * time.Minute)
    util.CE(err)
    _, _ = ctx.JSON(iris.Map{"status": http.StatusOK, "url": url})
}
