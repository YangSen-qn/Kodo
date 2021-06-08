package excel_cmd

import (
	"github.com/YangSen-qn/Kodo/cmd/common"
	"github.com/spf13/cobra"
)

type excelPerformer struct {
	config        *common.CommonPerformer
}

func NewExcelPerformerPerformer() *excelPerformer {
	return &excelPerformer{
		config: common.NewPerformer(),
	}
}

func ConfigCMD(superCMD *cobra.Command) {

	cmd := &cobra.Command{
		Use:     "excel",
		Short:   "excel cmd",
		Long:    "",
		Example: "",
	}
	configCopyCMD(cmd)
	superCMD.AddCommand(cmd)
}
