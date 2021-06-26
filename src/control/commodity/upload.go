package commodity

import (
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/control"
)

func (c *Commodity) Upload(ctx iris.Context) {
    control.Upload(ctx, &control.UploadProps{
        Type:   "image",
        Path:   c.cfg.S3Path.CommodityOrigin,
        Bucket: c.cfg.S3Bucket.Commodity,
        S3Sess: c.s3Sess,
    })
}
