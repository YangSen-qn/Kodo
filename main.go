package main

import (
	"fmt"

	"github.com/YangSen-qn/Kodo/cmd/root"
	"github.com/YangSen-qn/Kodo/cmd/uplog"
)

const _debug = false

func main() {
	if _debug {
		speed := uplog.NewSpeedPerformer()

		speed.Execute(nil, nil)

	} else {
		if err := root.LoadCMD(); err != nil {
			fmt.Println("error:", err)
		}
	}
}
