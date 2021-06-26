package commodity

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/control"
    "github.com/sia-fl/medicine/src/model"
    "github.com/sia-fl/medicine/src/util"
    "net/http"
)

type formQuery struct {
    control.FormQuery
    Name1     string   `url:"name1"`
    Name2     string   `url:"name2"`
    Name3     string   `url:"name3"`
    Cold      []int8   `url:"cold"`
    Herb      []int8   `url:"herb"`
    Danger    []int8   `url:"danger"`
    CreatedAt []string `url:"created_at"`
    UpdatedAt []string `url:"updated_at"`
}

func (c *Commodity) Query(ctx iris.Context) {
    var f formQuery
    util.CE(ctx.ReadQuery(&f))
    var m []model.Commodity
    q := control.Page(c.orm, f.ToPageFilter())
    if f.Name1 != "" {
        q.Where("name1 like ?", fmt.Sprintf("%%%s%%", f.Name1))
    }
    if f.Name2 != "" {
        q.Where("name2 like ?", fmt.Sprintf("%%%s%%", f.Name2))
    }
    if f.Name3 != "" {
        q.Where("name3 like ?", fmt.Sprintf("%%%s%%", f.Name3))
    }
    if len(f.CreatedAt) == 2 {
        q.Where("created_at > ?", f.CreatedAt[0])
        q.Where("created_at < ?", f.CreatedAt[1])
    }
    if len(f.UpdatedAt) == 2 {
        q.Where("updated_at > ?", f.UpdatedAt[0])
        q.Where("updated_at < ?", f.UpdatedAt[1])
    }
    if len(f.Herb) == 1 {
        q.Where("herb = ?", f.Herb)
    }
    if len(f.Cold) == 1 {
        q.Where("cold = ?", f.Cold)
    }
    if len(f.Danger) == 1 {
        q.Where("danger = ?", f.Danger)
    }
    if f.Sort != "" && f.SortDirection != "" {
        q = q.Order(fmt.Sprintf("%s %s", f.Sort, f.SortDirection))
    }
    q.Find(&m)
    var total int64
    c.orm.Model(&model.Commodity{}).Count(&total)
    _, _ = ctx.JSON(iris.Map{
        "status":    http.StatusOK,
        "data":      m,
        "total":     total,
        "page":      f.Page,
        "page_size": f.PageSize,
    })
}
