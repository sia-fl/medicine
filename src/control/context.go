package control

import (
    "context"
    "github.com/gorilla/sessions"
    "github.com/kataras/iris/v12"
)

type CtxMiddleware int

const (
    CtxSessions   CtxMiddleware = 996 + iota
    CtxSessionKey CtxMiddleware = iota
)

func WithSessions(s *sessions.CookieStore, sk string) func(iris.Context) {
    return func(ctx iris.Context) {
        _ctx1 := ctx.Request().Context()
        _ctx1 = context.WithValue(_ctx1, CtxSessions, s)
        _ctx1 = context.WithValue(_ctx1, CtxSessionKey, sk)
        ctx.ResetRequest(ctx.Request().WithContext(_ctx1))
        ctx.Next()
    }
}

func ContextFace(_ctx iris.Context, _k CtxMiddleware) interface{} {
    return _ctx.Request().Context().Value(_k)
}

func UId(ctx iris.Context) uint64 {
    sessFace := ContextFace(ctx, CtxSessions)
    sessKeyFace := ContextFace(ctx, CtxSessionKey)
    sess := sessFace.(*sessions.CookieStore)
    sessKey := sessKeyFace.(string)
    session, _ := sess.Get(ctx.Request(), sessKey)
    id, _ := session.Values["id"].(uint64)
    return id
}

func EId(ctx iris.Context) uint64 {
    sessFace := ContextFace(ctx, CtxSessions)
    sessKeyFace := ContextFace(ctx, CtxSessionKey)
    sess := sessFace.(*sessions.CookieStore)
    sessKey := sessKeyFace.(string)
    session, _ := sess.Get(ctx.Request(), sessKey)
    id, _ := session.Values["e_id"].(uint64)
    return id
}
