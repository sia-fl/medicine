package model

import "github.com/sia-fl/medicine/src/util"

type Category struct {
	util.Model
	Id     uint64 `json:"id"`
	EId    uint64 `json:"e_id"`
	PId    uint64 `json:"p_id"`
	Status uint8  `json:"status"`
	Name   string `json:"name"`
}
