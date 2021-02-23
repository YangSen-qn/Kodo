package version

import (
	"context"
	"fmt"
	"github.com/YangSen-qn/cmd"
)

var version = "v1.0.0"

var versionCmd = (&cmd.CommandBuilder {
	Use:   "version",
	Short: "show version",
	ExecuteFunction: func(context context.Context) error {
		fmt.Println(version)
		return nil
	},
}).Build()

func ConfigVersionCMD(superCMD *cmd.Command) {
	superCMD.AddCMD(versionCmd)
}
