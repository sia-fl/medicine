package main

import (
    "context"
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/cfg"
    controlCommodity "github.com/sia-fl/medicine/src/control/commodity"
    "github.com/sia-fl/medicine/src/inject"
    "github.com/sia-fl/medicine/src/util"
    "time"
)

func main() {
    container := inject.Inject()
    util.CE(container.Provide(func() *inject.CfgPath { return nil }))
    err := container.Invoke(func(
        cf *cfg.App,
        appAdmin inject.AppAdmin,
        controlCommodity *controlCommodity.Commodity,
    ) {
        iris.RegisterOnInterrupt(func() {
            ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
            _ = appAdmin.App.Shutdown(ctx)
            cancel()
        })
        if cf.AppAdminAddr != nil {
            go func() {
                inject.WithApiPartyCommodity(appAdmin.App.Party("/api"), controlCommodity)
                addr := fmt.Sprintf("%s:%s", cf.AppAdminAddr.Host, cf.AppAdminAddr.Port)
                err := appAdmin.App.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed))
                util.CE(err)
            }()
        }
    })
    util.CE(err)
    select {}
}
