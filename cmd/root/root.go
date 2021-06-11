package root

import (
	"github.com/YangSen-qn/Kodo/cmd/excel_cmd"
	"github.com/YangSen-qn/Kodo/cmd/info"
	"github.com/YangSen-qn/Kodo/cmd/ip"
	"github.com/YangSen-qn/Kodo/cmd/split"
	"github.com/YangSen-qn/Kodo/cmd/uplog"
	"github.com/YangSen-qn/Kodo/cmd/version"
	"github.com/spf13/cobra"
)

func LoadCMD() error {
	var rootCMD = &cobra.Command{
		Use:     "kodo",
		Short:   "my kodo command",
		Long:    "",
		Example: "",
		Version: version.Version,
	}

	version.ConfigCMD(rootCMD)
	info.ConfigCMD(rootCMD)
	uplog.ConfigCMD(rootCMD)
	uplog.ConfigTimeoutCMD(rootCMD)
	uplog.ConfigAllTypeCMD(rootCMD)
	uplog.ConfigSpeedCMD(rootCMD)
	uplog.ConfigSuccessRateCMD(rootCMD)
	ip.ConfigCMD(rootCMD)
	split.ConfigCMD(rootCMD)
	excel_cmd.ConfigCMD(rootCMD)
	return rootCMD.Execute()
}
