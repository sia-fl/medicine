package commodity

import (
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/model"
    "github.com/sia-fl/medicine/src/util"
    "net/http"
)

func (c *Commodity) Find(ctx iris.Context) {
    id, err := ctx.Params().GetUint64("id")
    util.CE(err)
    var m model.Commodity
    c.orm.Where("id = ?", id).First(&m)
    _, _ = ctx.JSON(iris.Map{"status": http.StatusOK, "data": m})
}
