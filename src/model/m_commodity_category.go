package model

import "github.com/sia-fl/medicine/src/util"

type CommodityCategory struct {
    util.Model
    Id     uint64 `gorm:"primaryKey;autoIncrement:false" json:"id"`
    EId    uint64 `gorm:"primaryKey;autoIncrement:false" json:"e_id"`
    PId    uint64 `json:"p_id"`
    Status uint8  `json:"status"`
    Name   string `json:"name"`
}

type CommodityCommodityCategory struct {
    CommodityId          uint64 `json:"commodity_id"`
    CommodityCategoryId  uint64 `json:"commodity_category_id"`
    CommodityCategoryEId uint64 `json:"commodity_category_e_id"`
}
