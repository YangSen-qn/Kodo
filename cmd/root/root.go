package root

import (
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
	ip.ConfigCMD(rootCMD)
	split.ConfigCMD(rootCMD)
	return rootCMD.Execute()
}
