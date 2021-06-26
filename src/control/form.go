package control

import "github.com/sia-fl/medicine/src/proto"

type FormDel struct {
    Ids []uint64 `json:"ids"`
}

type FormQuery struct {
    Page          uint32 `url:"page"`
    PageSize      uint32 `url:"page_size"`
    Sort          string `url:"sort"`
    SortDirection string `url:"sort_direction"`
}

func (f *FormQuery) ToPageFilter() *proto.PageFilter {
    return &proto.PageFilter{Page: f.Page, PageSize: f.PageSize}
}
