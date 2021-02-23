package query

import (
	"github.com/YangSen-qn/cmd"
)

var queryCMD = (&cmd.CommandBuilder{
	Use:             "query",
	Short:           "query something",
	Example:         "",
}).Build()


func ConfigQueryCMD(superCMD *cmd.Command) {
	configLogCMD(queryCMD)

	superCMD.AddCMD(queryCMD)
}

