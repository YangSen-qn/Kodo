package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/YangSen-qn/Kodo/cmd/root"
	"github.com/YangSen-qn/Kodo/core/log"
)

const _debug = true

func main() {
	if _debug {
		s := `[{"key":"version", "region":{"location":0,"length":2}}]`
		var typeParamList []*log.TypeParam
		if err := json.Unmarshal([]byte(s), &typeParamList); err != nil {
			fmt.Println(errors.New("type list err:" + err.Error()))
			return
		}

		fmt.Println("=== end ===")
	} else {
		if err := root.LoadCMD(); err != nil {
			fmt.Println("error:", err)
		}
	}
}
