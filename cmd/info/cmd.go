package info

import (
	"fmt"
	"strings"

	"github.com/YangSen-qn/Kodo/cmd/common"
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/config"
	"github.com/spf13/cobra"
)

type InfoPerformer struct {
	config *common.CommonPerformer
}

type info struct {
	AK         string   `json:"ak"`
	SK         string   `json:"sk"`
	BucketList []string `json:"bucketList"`
}

func (i *info) String() string {
	infoList := make([]string, 0, 2)
	infoList = append(infoList, fmt.Sprintf("ak=%s", i.AK))
	infoList = append(infoList, fmt.Sprintf("ak=%s", i.AK))
	infoList = append(infoList, fmt.Sprintf("bucketList=%s", i.BucketList))
	return strings.Join(infoList, " ")
}

func NewIPPerformer() *InfoPerformer {
	return &InfoPerformer{
		config: common.NewPerformer(),
	}
}

func ConfigCMD(superCMD *cobra.Command) {

	performer := NewIPPerformer()

	cmd := &cobra.Command{
		Use:   "info",
		Short: "get kodo info",
		Run:   performer.Execute,
	}

	performer.BindToCMD(cmd)
	performer.config.BindToCMD(cmd)
	superCMD.AddCommand(cmd)
}

func (performer *InfoPerformer) BindToCMD(cmd *cobra.Command) {
}

func (performer *InfoPerformer) Execute(cmd *cobra.Command, args []string) {
	performer.config.Execute(cmd, args)

	i := &info{
		AK:         config.AK,
		SK:         config.SK,
		BucketList: []string{
			"kodo-phone-zone0-space",
			"kodo-phone-zone1-space",
			"kodo-phone-zone2-space",
			"kodo-phone-as0-space",
			"kodo-phone-na0-space"},
	}
	output.I().Output(i)
}
