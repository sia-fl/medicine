package commodity

import (
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/control"
    "github.com/sia-fl/medicine/src/model"
    "github.com/sia-fl/medicine/src/util"
    "net/http"
)

func (c *Commodity) Del(ctx iris.Context) {
    var f control.FormDel
    util.CE(ctx.ReadJSON(&f))
    c.orm.Delete(&model.Commodity{}, f.Ids)
    _, _ = ctx.JSON(iris.Map{"status": http.StatusOK})
}
