package excel_cmd

import (
	"fmt"
	"github.com/YangSen-qn/Kodo/cmd/common"
	"github.com/YangSen-qn/Kodo/cmd/excel"
	"github.com/spf13/cobra"
)

type copyPerformer struct {
	config        *common.CommonPerformer
	fromFile      string
	fromSheet     string
	fromKeyColumn string
	fromColumn    string
	toFile        string
	toSheet       string
	toKeyColumn   string
	toColumn      string
	defaultValue  string
}

func NewCopyPerformer() *copyPerformer {
	return &copyPerformer{
		config: common.NewPerformer(),
	}
}

func configCopyCMD(superCMD *cobra.Command) {

	performer := NewCopyPerformer()

	cmd := &cobra.Command{
		Use:     "copy",
		Short:   "excel copy one column to another column, copy success while the value of fromKeyColumn is equal to toKeyColumn, not equal will set default value",
		Long:    "",
		Example: "",
		Run:     performer.Execute,
	}

	performer.BindLogCMDToPerformer(cmd)

	superCMD.AddCommand(cmd)
}

func (performer *copyPerformer) BindLogCMDToPerformer(command *cobra.Command) {
	command.Flags().StringVarP(&performer.fromFile, "from-file", "", "", "source file")
	command.Flags().StringVarP(&performer.fromSheet, "from-sheet", "", "", "source sheet")
	command.Flags().StringVarP(&performer.fromKeyColumn, "from-key-column", "", "A", "source key column, default A")
	command.Flags().StringVarP(&performer.fromColumn, "from-column", "", "", "source column")
	command.Flags().StringVarP(&performer.toFile, "to-file", "", "", "target file")
	command.Flags().StringVarP(&performer.toSheet, "to-sheet", "", "", "target sheet")
	command.Flags().StringVarP(&performer.toKeyColumn, "to-key-column", "", "A", "target key column, default A")
	command.Flags().StringVarP(&performer.toColumn, "to-column", "", "", "target column")
}

func (performer *copyPerformer) Execute(cmd *cobra.Command, args []string) {

	if performer.toFile == "" {
		return
	}

	if performer.toSheet == "" {
		return
	}

	if performer.toColumn == "" {
		return
	}

	if performer.fromFile == "" {
		return
	}

	if performer.fromSheet == "" {
		return
	}

	if performer.fromColumn == "" {
		return
	}

	sourceSheet, err := excel.NewSheet(performer.fromFile, performer.fromSheet)
	if err != nil {
		return
	}

	targetSheet, err := excel.NewSheet(performer.toFile, performer.toSheet)
	if err != nil {
		return
	}

	targetMaxRows, err := sourceSheet.MaxRows();
	if err != nil {
		return
	}

	fmt.Println("== targetMaxRows:", targetMaxRows)

	for targetRow := 1; targetRow < targetMaxRows; targetRow++{
		targetKeyRowValue, err := targetSheet.GetCellValue(performer.toKeyColumn, targetRow)
		if err != nil {
			fmt.Println("== targetKeyRowValue get err:", err)
			continue
		}
		fmt.Println("== targetKeyRowValue:", targetKeyRowValue)

		fromMaxRows, err := sourceSheet.MaxRows();
		if err != nil {
			return
		}

		var formValue interface{}
		for fromRow := 1; fromRow < fromMaxRows; fromRow++ {
			fromKeyRowValue, err := sourceSheet.GetCellValue(performer.fromKeyColumn, fromRow)
			if err != nil {
				continue
			}

			if fromKeyRowValue != targetKeyRowValue {
				continue
			}

			fmt.Println("== fromKeyRowValue:", fromKeyRowValue)

			value, err := sourceSheet.GetCellValue(performer.fromColumn, fromRow)
			if err != nil {
				continue
			}

			formValue = value

			break
		}

		fmt.Println("== fromValue:", formValue)

		if formValue != nil {
			fmt.Println("== copy targetRow:", targetRow, " value:", formValue)
			targetSheet.SetCellValue(performer.toColumn, targetRow, formValue)
		} else {
			fmt.Println("== copy targetRow:", targetRow, " value:", 0)
			targetSheet.SetCellValue(performer.toColumn, targetRow, 0)
		}
	}



	err = targetSheet.Save()
	if err != nil {
		fmt.Println("== save error:", err)
	}
}
