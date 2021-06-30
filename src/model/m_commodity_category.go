package model

import (
	"github.com/sia-fl/medicine/src/util"
	"gorm.io/gorm"
)

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

func (c CommodityCategory) Add(o *gorm.DB, f interface{}, id uint64, eId uint64) error {
	m := c
	util.ConvStruct(f, &m)
	m.Id = id
	m.EId = eId
	return o.Create(&m).Error
}

func (c CommodityCategory) Edit(o *gorm.DB, f interface{}) error {
	m := c
	util.ConvStruct(f, &m)
	return o.Updates(&m).Error
}

func (c CommodityCategory) Find(o *gorm.DB) *gorm.DB {
	return o.Model(c).Select("id", "p_id", "status", "name")
}

func (c CommodityCategory) Query(o *gorm.DB) *gorm.DB {
	return o.Model(c).Select("id", "p_id", "status", "name", "created_at", "updated_at")
}

func (c CommodityCategory) Delete(o *gorm.DB) error {
	return o.Delete(&c).Error
}
