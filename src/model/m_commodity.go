package model

import "github.com/sia-fl/medicine/src/util"

type Commodity struct {
    util.Model
    Id             uint64 `json:"id" gorm:"primaryKey"`
    Status         uint8  `json:"status"`
    ImageCover     string `json:"image_cover"`
    Name1          string `json:"name1"`
    Name2          string `json:"name2"`
    Name3          string `json:"name3"`
    Cold           uint8  `json:"cold"`
    Herb           uint8  `json:"herb"`
    Danger         uint8  `json:"danger"`
    Spec           string `json:"spec"`
    SpecUnit       string `json:"spec_unit"`
    TypeStore      string `json:"type_store"`
    ExpiryDateNum  int8   `json:"expiry_date_num"`
    ExpiryDateUnit string `json:"expiry_date_unit"`

    Dosages    []CommodityDosage   `json:"dosages" gorm:"many2many:commodity_commodity_dosages;"`
    Categories []CommodityCategory `json:"categories" gorm:"many2many:commodity_commodity_categories;"`
}
