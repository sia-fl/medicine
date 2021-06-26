package inject

import (
    "github.com/go-playground/validator/v10"
    "github.com/gorilla/sessions"
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/cfg"
    "github.com/sia-fl/medicine/src/control"
    controlCommodity "github.com/sia-fl/medicine/src/control/commodity"
    "github.com/sia-fl/medicine/src/util"
    "go.uber.org/dig"
)

type appDepend struct {
    dig.In
    Sess      *sessions.CookieStore
    Validator *validator.Validate
    Cfg       *cfg.App
}

const (
    SessionKeyAdmin string = "Admin"
)

func genAppAdmin(c *dig.Container) {
    util.CE(c.Provide(controlCommodity.NewControl))
    util.CE(c.Provide(genApp(SessionKeyAdmin), dig.Name("AppAdmin")))
}

func genApp(sessKey string) func(d appDepend) *iris.Application {
    return func(d appDepend) *iris.Application {
        app := iris.New()
        app.Validator = d.Validator
        app.Use(control.WithSessions(d.Sess, sessKey))
        if d.Cfg.Debug == true {
            app.Use(control.WithDebug())
        }
        return app
    }
}
