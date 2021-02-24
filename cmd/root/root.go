package root

import (
	"github.com/YangSen-qn/Kodo/cmd/uplog"
	"github.com/YangSen-qn/Kodo/cmd/version"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:        "kodo",
	Short:      "my kodo command",
	Long:       "",
	Example:    "",
	Version:    version.Version,
}


func LoadCMD() error {
	version.ConfigVersionCMD(rootCMD)
	uplog.ConfigLogCMD(rootCMD)
	return rootCMD.Execute()
}