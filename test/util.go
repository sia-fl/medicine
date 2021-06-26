package test

import (
    "github.com/sia-fl/medicine/src/inject"
)

func genCfgPath() *inject.CfgPath {
    return &inject.CfgPath{Path: "../cfg/used.toml"}
}

func Injection(f interface{}) {
    container := inject.Inject()
    _ = container.Provide(genCfgPath)
    err := container.Invoke(f)
    if err != nil {
        panic(err)
    }
}
