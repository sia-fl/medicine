package inject

import (
	"github.com/kataras/iris/v12"
	controlCommodityCategory "github.com/sia-fl/medicine/src/control/category"
	controlCommodity "github.com/sia-fl/medicine/src/control/commodity"
)

func WithApiPartyCommodity(
	a iris.Party,
	controlCommodity *controlCommodity.Commodity,
	controlCommodityCategory *controlCommodityCategory.Category,
) {
	partyCommodity := a.Party("/commodity")
	{
		partyCommodity.Post("/", controlCommodity.Add)
		partyCommodity.Delete("/", controlCommodity.Del)
		partyCommodity.Put("/{id:int}", controlCommodity.Edit)
		partyCommodity.Get("/{id:int}", controlCommodity.Find)
		partyCommodity.Get("/", controlCommodity.Query)
		partyCommodity.Post("/upload", controlCommodity.Upload)

		partyCommodityCategory := partyCommodity.Party("/category")
		{
			partyCommodityCategory.Post("/", controlCommodityCategory.Add)
			partyCommodityCategory.Delete("/", controlCommodityCategory.Del)
			partyCommodityCategory.Put("/{id:int}", controlCommodityCategory.Edit)
			partyCommodityCategory.Get("/{id:int}", controlCommodityCategory.Find)
			partyCommodityCategory.Get("/", controlCommodityCategory.Query)
		}
	}
}
