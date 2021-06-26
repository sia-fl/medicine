package inject

import (
    "github.com/kataras/iris/v12"
    controlCommodity "github.com/sia-fl/medicine/src/control/commodity"
)

func WithApiPartyCommodity(
    a iris.Party,
    controlCommodity *controlCommodity.Commodity,
) {
    partyCommodity := a.Party("/commodity")
    {
        partyCommodity.Post("/", controlCommodity.Add)
        partyCommodity.Delete("/", controlCommodity.Del)
        partyCommodity.Put("/{id:int}", controlCommodity.Edit)
        partyCommodity.Get("/{id:int}", controlCommodity.Find)
        partyCommodity.Get("/", controlCommodity.Query)
        partyCommodity.Post("/upload", controlCommodity.Upload)
    }
}
