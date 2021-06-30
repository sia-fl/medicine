package control

import (
	"fmt"
	"gorm.io/gorm"
)

type FormDel struct {
	Ids []uint64 `json:"ids"`
}

type FormPage struct {
	Page          uint32 `json:"page" url:"page"`
	PageSize      uint32 `json:"page_size" url:"page_size"`
	Sort          string `json:"sort" url:"sort"`
	SortDirection string `json:"sort_direction" url:"sort_direction"`
}

type FormAt struct {
	CreatedAt []string `json:"created_at" url:"created_at"`
	UpdatedAt []string `json:"updated_at" url:"updated_at"`
}

func (f *FormPage) WithPage(q *gorm.DB) *gorm.DB {
	if f.Sort != "" && f.SortDirection != "" {
		q = q.Order(fmt.Sprintf("%s %s", f.Sort, f.SortDirection))
	}
	return q.Offset(int((f.Page - 1) * f.PageSize)).Limit(int(f.PageSize))
}

func (f FormAt) WithAt(q *gorm.DB) *gorm.DB {
	if len(f.CreatedAt) == 2 {
		q.Where("created_at > ?", f.CreatedAt[0])
		q.Where("created_at < ?", f.CreatedAt[1])
	}
	if len(f.UpdatedAt) == 2 {
		q.Where("updated_at > ?", f.UpdatedAt[0])
		q.Where("updated_at < ?", f.UpdatedAt[1])
	}
	return q
}
