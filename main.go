package main

import (
	"fmt"
	"github.com/YangSen-qn/Kodo/cmd/root"
)

func main() {

	if err := root.LoadCMD(); err != nil {
		fmt.Println("error:", err)
	}
}
