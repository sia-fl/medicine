package main

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sia-fl/medicine/cfg"
	controlCommodityCategory "github.com/sia-fl/medicine/src/control/category"
	controlCommodity "github.com/sia-fl/medicine/src/control/commodity"
	"github.com/sia-fl/medicine/src/inject"
	"github.com/sia-fl/medicine/src/util"
	"go.uber.org/dig"
	"time"
)

type InvokeCol struct {
	dig.In
	Cf                       *cfg.App
	AppAdmin                 *iris.Application `name:"AppAdmin"`
	ControlCommodity         *controlCommodity.Commodity
	ControlCommodityCategory *controlCommodityCategory.Category `name:"CommodityCategories"`
}

func main() {
	container := inject.Inject()
	util.CE(container.Provide(func() *inject.CfgPath { return nil }))
	err := container.Invoke(func(c InvokeCol) {
		iris.RegisterOnInterrupt(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			_ = c.AppAdmin.Shutdown(ctx)
			cancel()
		})
		if c.Cf.AppAdminAddr != nil {
			go func() {
				inject.WithApiPartyCommodity(
					c.AppAdmin.Party("/api"),
					c.ControlCommodity,
					c.ControlCommodityCategory,
				)
				addr := fmt.Sprintf("%s:%s", c.Cf.AppAdminAddr.Host, c.Cf.AppAdminAddr.Port)
				err := c.AppAdmin.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed))
				util.CE(err)
			}()
		}
	})
	util.CE(err)
	select {}
}
