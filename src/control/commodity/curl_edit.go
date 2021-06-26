package commodity

import (
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/model"
    "github.com/sia-fl/medicine/src/util"
    "gorm.io/gorm"
)

type editForm struct {
    Id uint64
    form
}

func (c *Commodity) Edit(ctx iris.Context) {
    var f editForm
    var m model.Commodity
    util.CE(ctx.ReadJSON(&f))
    f.form.ToModel(&m)
    id, err := ctx.Params().GetUint64("id")
    util.CE(err)
    util.CE(c.orm.Transaction(func(tx *gorm.DB) error {
        err := tx.Where("commodity_id = ?", id).Delete(&model.CommodityCommodityCategory{}).Error
        if err != nil {
            return err
        }
        err = tx.Where("commodity_id = ?", id).Delete(&model.CommodityCommodityDosage{}).Error
        if err != nil {
            return err
        }
        tx.Model(&model.Commodity{Id: id}).Updates(m)
        return nil
    }))
}
