package category

import (
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/src/control"
	"github.com/sia-fl/medicine/src/util"
	"net/http"
)

func (c *Category) Find(ctx iris.Context) {
	id, err := ctx.Params().GetUint64("id")
	util.CE(err)
	result := make(map[string]interface{})
	q := c.M().Find(c.orm)
	q = q.Where("e_id = ?", control.EId(ctx))
	q = q.Where("id = ?", id)
	util.CE(q.First(&result).Error)
	_, _ = ctx.JSON(iris.Map{"status": http.StatusOK, "data": result})
}
