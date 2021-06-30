package category

import (
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/src/control"
	"github.com/sia-fl/medicine/src/util"
	"net/http"
)

func (c *Category) Del(ctx iris.Context) {
	var f control.FormDel
	util.CE(ctx.ReadJSON(&f))
	q := c.orm
	q = q.Where("e_id = ?", control.EId(ctx))
	q = q.Where("id in ?", f.Ids)
	util.CE(c.M().Delete(q))
	_, _ = ctx.JSON(iris.Map{"status": http.StatusOK})
}
