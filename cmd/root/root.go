package root

import (
	"github.com/YangSen-qn/Kodo/cmd/split"
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
	version.ConfigCMD(rootCMD)
	uplog.ConfigCMD(rootCMD)
	split.ConfigCMD(rootCMD)
	return rootCMD.Execute()
}
