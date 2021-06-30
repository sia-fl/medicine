package category

import (
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/src/util"
	"net/http"
)

func (c *Category) Query(ctx iris.Context) {
	f := c.formQuery.(InterfaceFormQuery).ReadQuery(ctx)
	result := make([]map[string]interface{}, 20)
	q := c.M().Query(c.orm)
	q = f.WithQuery(q)
	util.CE(q.Find(&result).Error)
	var total int64
	c.orm.Model(c.model).Count(&total)
	formPage := f.GetFormPage()
	_, _ = ctx.JSON(iris.Map{
		"status":    http.StatusOK,
		"data":      result,
		"total":     total,
		"page":      formPage.Page,
		"page_size": formPage.PageSize,
	})
}
