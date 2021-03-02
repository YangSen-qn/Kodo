package common

import (
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/spf13/cobra"
)

type CommonPerformer struct {
	OutputFormat string
}

func NewPerformer() *CommonPerformer {
	return &CommonPerformer{OutputFormat: "human"}
}

func (performer *CommonPerformer) BindToCMD(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&performer.OutputFormat, "output-format", "", "", "format that output, eg:human, json")
}

func (performer *CommonPerformer) Execute(cmd *cobra.Command, args []string) {
	if performer.OutputFormat == "json" {
		output.SetOutput(&output.JsonOutput{true})
	} else {
		output.SetOutput(&output.DefaultOutput{IsColorful: true})
	}
}
