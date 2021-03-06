package test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/sia-fl/medicine/src/control"
	controlCommodityCategory "github.com/sia-fl/medicine/src/control/category"
	controlCommodity "github.com/sia-fl/medicine/src/control/commodity"
	"github.com/sia-fl/medicine/src/inject"
	"github.com/sia-fl/medicine/src/model"
	"github.com/sia-fl/medicine/src/proto"
	"github.com/sia-fl/medicine/src/util"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

var (
	exampleSpecUnit       = []string{"gram", "box", "bottle"}
	exampleExpiryDateUnit = []string{"day", "week", "month", "year"}
	exampleTypeStore      = []string{"normal", "shady", "refrigeration", "freezing"}
)

func TestCommodity(t *testing.T) {
	t.Run("testDatabaseCommodity", testDatabaseCommodity)
	t.Run("testFakeCommodity", testFakeCommodity)
}

func testDatabaseCommodity(_ *testing.T) {
	Injection(func(
		orm *gorm.DB,
	) {
		util.CE(orm.Migrator().DropTable(&model.CommodityCommodityCategory{}))
		util.CE(orm.Migrator().DropTable(&model.CommodityCommodityDosage{}))
		util.CE(orm.Migrator().DropTable(&model.CommodityCategory{}))
		util.CE(orm.Migrator().DropTable(&model.CommodityDosage{}))
		util.CE(orm.Migrator().DropTable(&model.Commodity{}))
		util.CE(orm.Migrator().CreateTable(&model.CommodityCommodityCategory{}))
		util.CE(orm.Migrator().CreateTable(&model.CommodityCommodityDosage{}))
		util.CE(orm.Migrator().CreateTable(&model.CommodityCategory{}))
		util.CE(orm.Migrator().CreateTable(&model.CommodityDosage{}))
		var i interface{}
		i = &model.Commodity{}
		util.CE(orm.Migrator().CreateTable(i))
	})
}

func testFakeCommodity(_ *testing.T) {
	Injection(func(
		orm *gorm.DB,
		idsRpcSer proto.IdInMapServicesClient,
	) {
		length := uint64(30)
		var commodityItem []model.Commodity
		for i := length - length; i < length; i++ {
			commodityItem = append(commodityItem, model.Commodity{
				Status:         uint8(gofakeit.Number(1, 2)),
				Name1:          gofakeit.Name(),
				Name2:          gofakeit.Name(),
				Name3:          gofakeit.Name(),
				Cold:           uint8(gofakeit.Number(1, 2)),
				Herb:           uint8(gofakeit.Number(1, 2)),
				Danger:         uint8(gofakeit.Number(1, 2)),
				SpecUnit:       exampleSpecUnit[gofakeit.Number(0, 2)],
				Spec:           gofakeit.DomainName(),
				TypeStore:      exampleTypeStore[gofakeit.Number(0, 3)],
				ExpiryDateUnit: exampleExpiryDateUnit[gofakeit.Number(0, 2)],
				ExpiryDateNum:  int8(gofakeit.Number(10, 30)),
			})
		}
		var dosageItem []model.CommodityDosage
		tmpIdRes, err := idsRpcSer.GetIdArrInMap(context.Background(), &proto.GetIdArrInMapReq{EId: 0, Length: length, Name: "commodity_dosages"})
		util.CE(err)
		for i := tmpIdRes.Id - length; i < tmpIdRes.Id; i++ {
			gofakeit.Seed(0)
			tmpDosage := model.CommodityDosage{
				EId:  0,
				PId:  uint64(gofakeit.Number(0, int(i))),
				Name: "??????_" + strconv.FormatUint(i, 10),
			}
			tmpDosage.Id = i
			dosageItem = append(dosageItem, tmpDosage)
		}
		var categoryItem []model.CommodityCategory
		tmpIdRes, err = idsRpcSer.GetIdArrInMap(context.Background(), &proto.GetIdArrInMapReq{EId: 0, Length: length, Name: "commodity_categories"})
		util.CE(err)
		for i := tmpIdRes.Id - length; i < tmpIdRes.Id; i++ {
			gofakeit.Seed(0)
			tmpCategory := model.CommodityCategory{
				EId:  0,
				PId:  uint64(gofakeit.Number(0, int(i))),
				Name: "????????????_" + strconv.FormatUint(i, 10),
			}
			tmpCategory.Id = i
			categoryItem = append(categoryItem, tmpCategory)
		}
		orm.Create(&commodityItem)
		orm.Create(&dosageItem)
		orm.Create(&categoryItem)
	})
}

func MaxCategoryId(o *gorm.DB) uint64 {
	row := o.Model(&model.CommodityCategory{}).Select("max(id) max").Row()
	var max uint64
	_ = row.Scan(&max)
	return max
}

func MaxDosageId(o *gorm.DB) uint64 {
	row := o.Model(&model.CommodityDosage{}).Select("max(id) max").Row()
	var max uint64
	_ = row.Scan(&max)
	return max
}

type InvokeColTestCurlCommodity struct {
	dig.In
	Orm                      *gorm.DB
	App                      *iris.Application `name:"AppAdmin"`
	ControlCommodity         *controlCommodity.Commodity
	ControlCommodityCategory *controlCommodityCategory.Category `name:"CommodityCategories"`
}

func TestCurlCommodity(t *testing.T) {
	Injection(func(c InvokeColTestCurlCommodity) {
		c.App.Use(control.WithDebug())
		inject.WithApiPartyCommodity(
			c.App.Party("/api"),
			c.ControlCommodity,
			c.ControlCommodityCategory,
		)
		guess := httptest.New(t, c.App)
		categoryId := MaxCategoryId(c.Orm)

		// Add Category
		commodityCategoryBodyAdd := iris.Map{
			"name":   "????????????",
			"p_id":   categoryId,
			"status": 1,
		}
		guess.POST("/api/commodity/category").WithJSON(commodityCategoryBodyAdd).Expect().Status(200)

		// ADD ID
		row := c.Orm.Model(&model.CommodityCategory{}).Where("e_id", 1).Select("max(id) max").Row()
		var max string
		_ = row.Scan(&max)

		// Edit Category
		commodityCategoryBodyEdit := iris.Map{
			"name":   "????????????",
			"p_id":   categoryId - 1,
			"status": 2,
		}
		guess.PUT("/api/commodity/category/" + max).WithJSON(commodityCategoryBodyEdit).Expect().Status(200)

		// FIND ONE
		guess.GET("/api/commodity/category/" + max).Expect().Status(200)

		// FIND MANY
		commodityCategoryBodyQuery := iris.Map{
			"page":      1,
			"page_size": 2,
			"name":      "??????",
		}
		guess.GET("/api/commodity/category").WithQueryObject(commodityCategoryBodyQuery).Expect().Status(200)

		// MIN Category ID
		row = c.Orm.Model(&model.CommodityCategory{}).Where("e_id", 1).Select("min(id) min").Row()
		var min uint64
		_ = row.Scan(&min)

		// Del Category
		guess.DELETE("/api/commodity/category").WithJSON(iris.Map{"ids": []uint64{min}}).Expect().Status(200)

		// ADD
		dosageId := MaxDosageId(c.Orm)
		commodityBodyAdd := iris.Map{
			"category_ids": []uint64{categoryId},
			"dosage_ids":   []uint64{dosageId},
			"image_cover_arr": []string{
				"12345678123456781234567812345678.jpg",
				"87654321876543218765432187654321.jpg",
			},
			"status":           1,
			"name_1":           "????????????",
			"name_2":           "????????????",
			"name_3":           "??????????????????",
			"spec_unit":        exampleSpecUnit[gofakeit.Number(0, 2)],
			"spec":             gofakeit.DomainName(),
			"expiry_date_unit": exampleExpiryDateUnit[gofakeit.Number(0, 3)],
			"expiry_date_num":  gofakeit.Number(10, 30),
			"type_store":       "normal",
		}
		guess.POST("/api/commodity").WithJSON(commodityBodyAdd).Expect().Status(200)

		// ADD ID
		row = c.Orm.Model(&model.Commodity{}).Select("max(id) max").Row()
		_ = row.Scan(&max)

		// EDIT
		commodityBodyEdit := iris.Map{
			"category_ids":     []uint64{categoryId - 10},
			"dosage_ids":       []uint64{dosageId - 10},
			"status":           2,
			"name1":            "????????????update",
			"name2":            "????????????update",
			"name3":            "??????????????????update",
			"spec_unit":        exampleSpecUnit[gofakeit.Number(0, 2)],
			"spec":             gofakeit.DomainName(),
			"expiry_date_unit": exampleExpiryDateUnit[gofakeit.Number(0, 3)],
			"expiry_date_num":  gofakeit.Number(10, 30),
			"type_store":       "normal",
		}
		guess.PUT("/api/commodity/" + max).WithJSON(commodityBodyEdit).Expect().Status(200)

		// FIND ONE
		guess.GET("/api/commodity/" + max).Expect().Status(200)

		// FIND MANY
		commodityBodyQuery := iris.Map{
			"page":      1,
			"page_size": 2,
			"name1":     "update",
		}
		guess.GET("/api/commodity").WithQueryObject(commodityBodyQuery).Expect().Status(200)

		// MIN ID
		row = c.Orm.Model(&model.Commodity{}).Select("min(id) min").Row()
		_ = row.Scan(&min)

		// DEL
		guess.DELETE("/api/commodity").WithJSON(iris.Map{"ids": []uint64{min}}).Expect().Status(200)

		// Upload
		uploadCommodityBody := map[string]interface{}{
			"file_md_5": "12345678123456781234567812345678",
			"file_ext":  "jpeg",
		}
		guess.POST("/api/commodity/upload").WithJSON(uploadCommodityBody).Expect().Status(200)
	})
}
