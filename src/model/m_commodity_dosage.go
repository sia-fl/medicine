package model

import "github.com/sia-fl/medicine/src/util"

type CommodityDosage struct {
    util.Model
    Id     uint64 `gorm:"primaryKey;autoIncrement:false" json:"id"`
    EId    uint64 `gorm:"primaryKey;autoIncrement:false" json:"e_id"`
    PId    uint64 `json:"p_id"`
    Status uint8  `json:"status"`
    Name   string `json:"name"`
}

type CommodityCommodityDosage struct {
    CommodityId        uint64 `json:"commodity_id"`
    CommodityDosageId  uint64 `json:"commodity_dosage_id"`
    CommodityDosageEId uint64 `json:"commodity_dosage_e_id"`
}
