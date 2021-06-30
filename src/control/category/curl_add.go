package category

import (
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/src/control"
	"github.com/sia-fl/medicine/src/util"
	"net/http"
)

func (c *Category) Add(ctx iris.Context) {
	f := c.form
	util.CE(ctx.ReadJSON(&f))
	eId := control.EId(ctx)
	id := c.GetIdInMap(ctx.Request().Context(), eId)
	util.CE(c.M().Add(c.orm, f, id, control.EId(ctx)))
	_, _ = ctx.JSON(iris.Map{"status": http.StatusOK})
}
