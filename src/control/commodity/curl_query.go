package commodity

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/src/control"
	"github.com/sia-fl/medicine/src/model"
	"github.com/sia-fl/medicine/src/util"
	"net/http"
)

type formQuery struct {
	control.FormPage
	control.FormAt
	Name1  string `url:"name1"`
	Name2  string `url:"name2"`
	Name3  string `url:"name3"`
	Cold   []int8 `url:"cold"`
	Herb   []int8 `url:"herb"`
	Danger []int8 `url:"danger"`
}

func (c *Commodity) Query(ctx iris.Context) {
	var f formQuery
	util.CE(ctx.ReadQuery(&f))
	var m []model.Commodity
	q := c.orm
	q = f.WithAt(q)
	q = f.WithPage(q)
	if f.Name1 != "" {
		q.Where("name1 like ?", fmt.Sprintf("%%%s%%", f.Name1))
	}
	if f.Name2 != "" {
		q.Where("name2 like ?", fmt.Sprintf("%%%s%%", f.Name2))
	}
	if f.Name3 != "" {
		q.Where("name3 like ?", fmt.Sprintf("%%%s%%", f.Name3))
	}
	if len(f.Herb) == 1 {
		q.Where("herb = ?", f.Herb)
	}
	if len(f.Cold) == 1 {
		q.Where("cold = ?", f.Cold)
	}
	if len(f.Danger) == 1 {
		q.Where("danger = ?", f.Danger)
	}
	util.CE(q.Find(&m).Error)
	var total int64
	c.orm.Model(&model.Commodity{}).Count(&total)
	_, _ = ctx.JSON(iris.Map{
		"status":    http.StatusOK,
		"data":      m,
		"total":     total,
		"page":      f.Page,
		"page_size": f.PageSize,
	})
}
