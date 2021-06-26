package commodity

import (
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/model"
    "github.com/sia-fl/medicine/src/util"
    "net/http"
)

func (c *Commodity) Add(ctx iris.Context) {
    var f form
    var m model.Commodity
    util.CE(ctx.ReadJSON(&f))
    f.ToModel(&m)
    c.orm.Save(&m)
    _, _ = ctx.JSON(iris.Map{"status": http.StatusOK})
}
