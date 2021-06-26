package commodity

import (
    "context"
    "encoding/json"
    s3Session "github.com/aws/aws-sdk-go/aws/session"
    "github.com/go-playground/validator/v10"
    "github.com/gorilla/sessions"
    "github.com/sia-fl/medicine/cfg"
    "github.com/sia-fl/medicine/src/model"
    "github.com/sia-fl/medicine/src/proto"
    "github.com/sia-fl/medicine/src/util"
    "gorm.io/gorm"
)

type Commodity struct {
    orm        *gorm.DB
    cfg        *cfg.App
    sess       *sessions.CookieStore
    s3Sess     *s3Session.Session
    validate   *validator.Validate
    serId      proto.IdServicesClient
    serIdInMap proto.IdInMapServicesClient
}

type form struct {
    CategoryIds    []uint64 `json:"category_ids"`
    DosageIds      []uint64 `json:"dosage_ids"`
    ImageCoverArr  []string `json:"image_cover_arr"`
    Status         uint8    `json:"status"`
    Name1          string   `json:"name1"`
    Name2          string   `json:"name2"`
    Name3          string   `json:"name3"`
    Cold           uint8    `json:"cold"`
    Herb           uint8    `json:"herb"`
    Danger         uint8    `json:"danger"`
    Spec           string   `json:"spec"`
    SpecUnit       string   `json:"spec_unit"`
    TypeStore      string   `json:"type_store"`
    ExpiryDateNum  int8     `json:"expiry_date_num"`
    ExpiryDateUnit string   `json:"expiry_date_unit"`
}

func (f *form) ToModel(m *model.Commodity) {
    m.Status = f.Status
    if len(f.ImageCoverArr) > 0 {
        imageCover, err := json.Marshal(f.ImageCoverArr)
        util.CE(err)
        m.ImageCover = string(imageCover)
    }
    m.Name1 = f.Name1
    m.Name2 = f.Name2
    m.Name3 = f.Name3
    m.Cold = f.Cold
    m.Herb = f.Herb
    m.Danger = f.Danger
    m.Spec = f.Spec
    m.SpecUnit = f.SpecUnit
    m.TypeStore = f.TypeStore
    m.ExpiryDateNum = f.ExpiryDateNum
    m.ExpiryDateUnit = f.ExpiryDateUnit
    if m.Categories == nil {
        m.Categories = []model.CommodityCategory{}
    }
    if m.Dosages == nil {
        m.Dosages = []model.CommodityDosage{}
    }
    for _, id := range f.CategoryIds {
        commodityCategory := model.CommodityCategory{}
        commodityCategory.Id = id
        m.Categories = append(m.Categories, commodityCategory)
    }
    for _, id := range f.DosageIds {
        commodityDosage := model.CommodityDosage{}
        commodityDosage.Id = id
        m.Dosages = append(m.Dosages, commodityDosage)
    }
}

func (c *Commodity) GetIdInMap(eId uint64) uint64 {
    res, err := c.serIdInMap.GetIdInMap(context.Background(), &proto.GetIdInMapReq{
        EId:  eId,
        Name: "commodities",
    })
    util.CE(err)
    return res.Id
}

func (c *Commodity) GetId() uint64 {
    res, err := c.serId.GetId(context.Background(), &proto.GetIdReq{Name: "commodities"})
    util.CE(err)
    return res.Id
}

func NewControl(
    orm *gorm.DB,
    cfg *cfg.App,
    sess *sessions.CookieStore,
    s3Sess *s3Session.Session,
    validate *validator.Validate,
    serId proto.IdServicesClient,
    serIdInMap proto.IdInMapServicesClient,
) *Commodity {
    return &Commodity{
        cfg:        cfg,
        orm:        orm,
        sess:       sess,
        s3Sess:     s3Sess,
        validate:   validate,
        serId:      serId,
        serIdInMap: serIdInMap,
    }
}
