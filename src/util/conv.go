package util

import (
	"encoding/json"
)

func ConvStruct(b interface{}, a interface{}) {
	n, _ := json.Marshal(b)
	_ = json.Unmarshal(n, a)
}
