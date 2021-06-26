package control

import (
    "github.com/gorilla/sessions"
    "github.com/kataras/iris/v12"
    "github.com/sia-fl/medicine/src/util"
)

type UserInfoState struct {
    Id   uint64
    EId  uint64
    Name string
}

var UserInfo = &UserInfoState{
    Id:   1,
    EId:  1,
    Name: "全程高能",
}

func WithDebug() func(ctx iris.Context) {
    return func(ctx iris.Context) {
        sessFace := ContextFace(ctx, CtxSessions)
        sessKeyFace := ContextFace(ctx, CtxSessionKey)
        sess := sessFace.(*sessions.CookieStore)
        sessKey := sessKeyFace.(string)
        session, err := sess.Get(ctx.Request(), sessKey)
        util.CE(err)
        session.Values["id"] = UserInfo.Id
        session.Values["e_id"] = UserInfo.EId
        ctx.Next()
    }
}
