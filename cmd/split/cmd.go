package split

import (
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/util"
	"github.com/spf13/cobra"
)

type SplitPerformer struct {
	filePath  string
	separator string
}

func NewSplitPerformer() *SplitPerformer {
	return &SplitPerformer{}
}

func ConfigCMD(superCMD *cobra.Command) {

	performer := NewSplitPerformer()

	cmd := &cobra.Command{
		Use:   "split",
		Short: "split string",
		Run:   performer.Execute,
	}

	performer.BindToCMD(cmd)
	superCMD.AddCommand(cmd)
}

func (performer *SplitPerformer) BindToCMD(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&performer.filePath, "file", "f", "", "cut string from file")
	cmd.Flags().StringVarP(&performer.separator, "separator", "s", "", "the separator string that cut by")
}

func (performer *SplitPerformer) Execute(cmd *cobra.Command, args []string) {
	if len(performer.separator) == 0 {
		output.W().OutputFormat("the separator string can not be empty\n")
		return
	}

	if performer.filePath != "" {
		errorResult := make(chan error)
		partResult := make(chan string, 2)
		go util.SplitFromFile(performer.filePath, performer.separator, partResult, errorResult)

		complete := false
		for !complete {
			select {
			case <-errorResult:
				complete = true
				break
			case s, ok := <-partResult:
				if !ok {
					complete = true
					break
				}
				output.I().OutputFormat("%s\n", s)
			default:
			}
		}
	} else if len(args) > 0 {
		result := util.Split(args[0], performer.separator)
		for _, s := range result {
			output.I().OutputFormat("%s\n", s)
		}
	} else {
		output.W().OutputFormat("the string that cut is empty\n")
	}
}
