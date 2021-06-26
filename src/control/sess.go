package control

import (
    "github.com/gorilla/sessions"
    "net/http"
)

type Sess struct {
    s   *sessions.CookieStore
    key string
}

func (s *Sess) Get(r *http.Request) *sessions.Session {
    sess, err := s.s.Get(r, s.key)
    if err != nil {
        panic(err)
    }
    return sess
}
