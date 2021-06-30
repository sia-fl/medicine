package category

import (
	"context"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/cfg"
	"github.com/sia-fl/medicine/src/control"
	"github.com/sia-fl/medicine/src/proto"
	"github.com/sia-fl/medicine/src/util"
	"gorm.io/gorm"
)

type Category struct {
	tableName  string
	form       interface{}
	formQuery  interface{}
	model      interface{}
	orm        *gorm.DB
	sess       *sessions.CookieStore
	serIdInMap proto.IdInMapServicesClient
}

type InterfaceForm interface {
	ToModel(iris.Context, interface{})
}

type InterfaceFormQuery interface {
	ReadQuery(ctx iris.Context) InterfaceFormQuery
	WithQuery(q *gorm.DB) *gorm.DB
	GetFormPage() control.FormPage
}

type InterfaceModel interface {
	Add(orm *gorm.DB, form interface{}, id uint64, eId uint64) error
	Edit(orm *gorm.DB, form interface{}) error
	Find(orm *gorm.DB) *gorm.DB
	Query(orm *gorm.DB) *gorm.DB
	Delete(orm *gorm.DB) error
}

type Form struct {
	PId    uint64 `json:"p_id"`
	Status uint8  `json:"status"`
	Name   string `json:"name"`
}

type FormQuery struct {
	control.FormPage
	control.FormAt
	Name string `json:"name" url:"name"`
}

func (f FormQuery) WithQuery(q *gorm.DB) *gorm.DB {
	if f.Name != "" {
		q.Where("name like ?", fmt.Sprintf("%%%s%%", f.Name))
	}
	return f.WithAt(f.WithPage(q))
}

func (f FormQuery) GetFormPage() control.FormPage {
	return f.FormPage
}

func (f FormQuery) ReadQuery(ctx iris.Context) InterfaceFormQuery {
	_f1 := f
	util.CE(ctx.ReadQuery(&_f1))
	return _f1
}

func (c *Category) GetIdInMap(ctx context.Context, eId uint64) uint64 {
	res, err := c.serIdInMap.GetIdInMap(ctx, &proto.GetIdInMapReq{
		EId:  eId,
		Name: c.tableName,
	})
	util.CE(err)
	return res.Id
}

func NewControl(
	name string,
	form interface{},
	formQuery interface{},
	model interface{},
) func(
	orm *gorm.DB,
	dbCfg *cfg.Database,
	sess *sessions.CookieStore,
	serIdInMap proto.IdInMapServicesClient,
) *Category {
	return func(
		orm *gorm.DB,
		dbCfg *cfg.Database,
		sess *sessions.CookieStore,
		serIdInMap proto.IdInMapServicesClient,
	) *Category {
		return &Category{
			orm:        orm,
			tableName:  name,
			form:       form,
			formQuery:  formQuery,
			model:      model,
			sess:       sess,
			serIdInMap: serIdInMap,
		}
	}
}

func (c *Category) M() InterfaceModel {
	return c.model.(InterfaceModel)
}
