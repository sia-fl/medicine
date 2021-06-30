package category

import (
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/src/control"
	"github.com/sia-fl/medicine/src/util"
	"net/http"
)

func (c *Category) Edit(ctx iris.Context) {
	f := c.form
	util.CE(ctx.ReadJSON(&f))
	id, err := ctx.Params().GetUint64("id")
	util.CE(err)
	eId := control.EId(ctx)
	util.CE(c.M().Edit(c.orm.Where("e_id = ?", eId).Where("id = ?", id), f))
	_, _ = ctx.JSON(iris.Map{"status": http.StatusOK})
}
