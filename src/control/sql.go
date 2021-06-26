package control

import (
    "github.com/sia-fl/medicine/src/proto"
    "gorm.io/gorm"
)

func Page(o *gorm.DB, p *proto.PageFilter) *gorm.DB {
    return o.Offset(int((p.Page - 1) * p.PageSize)).Limit(int(p.PageSize))
}