package ip

import (
	"github.com/YangSen-qn/Kodo/cmd/common"
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/util"
	"github.com/spf13/cobra"
)

type IPPerformer struct {
	config *common.CommonPerformer
}

func NewIPPerformer() *IPPerformer {
	return &IPPerformer{
		config: common.NewPerformer(),
	}
}

func ConfigCMD(superCMD *cobra.Command) {

	performer := NewIPPerformer()

	cmd := &cobra.Command{
		Use:   "ip",
		Short: "get position by ip",
		Run:   performer.Execute,
	}

	performer.config.BindToCMD(cmd)
	superCMD.AddCommand(cmd)
}

func (performer *IPPerformer) Execute(cmd *cobra.Command, args []string) {
	performer.config.Execute(cmd, args)

	if len(args) == 0 {
		output.W().OutputFormat("the ip can not be empty\n")
		return
	}

	info , err := util.GetPositionByIP(args[0])
	if err != nil {
		output.E(err)
		return
	}

	output.I().Output(info)
}