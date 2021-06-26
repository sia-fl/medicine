package util

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/kataras/iris/v12"
)

type validationError struct {
    ActualTag string `json:"tag"`
    Namespace string `json:"namespace"`
    Kind      string `json:"kind"`
    Type      string `json:"type"`
    Value     string `json:"value"`
    Param     string `json:"param"`
}

func CE(errArr ...interface{}) {
    for i := 0; i < len(errArr); i++ {
        if err, ok := errArr[i].(error); ok {
            panic(err)
        }
    }
}

func ErrReType(ctx iris.Context) {
    defer func() {
        if err := recover(); err != nil {
            if errs, ok := err.(validator.ValidationErrors); ok {
                validationErrors := make([]validationError, 0, len(errs))
                for _, validationErr := range errs {
                    validationErrors = append(validationErrors, validationError{
                        ActualTag: validationErr.ActualTag(),
                        Namespace: validationErr.Namespace(),
                        Kind:      validationErr.Kind().String(),
                        Type:      validationErr.Type().String(),
                        Value:     fmt.Sprintf("%v", validationErr.Value()),
                        Param:     validationErr.Param(),
                    })
                }
                ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Key("errors", validationErrors))
                return
            }
        }
    }()
    ctx.Next()
}
