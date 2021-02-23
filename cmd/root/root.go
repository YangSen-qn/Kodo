package root

import (
	"github.com/YangSen-qn/Kodo/cmd/query"
	"github.com/YangSen-qn/Kodo/cmd/version"
	"github.com/YangSen-qn/cmd"
)

var rootCMD = (&cmd.CommandBuilder{
	Use:                    "cmd",
	Short:                  "just a demo",
	Version:                "0.0.1",
	BashCompletionFulsnction: "",
}).Build()

func init() {
	query.ConfigQueryCMD(rootCMD)
	version.ConfigVersionCMD(rootCMD)
}

func loadCMD() error {
	return rootCMD.Run()
}