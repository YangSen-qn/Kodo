package uplog

import (
	"github.com/YangSen-qn/Kodo/cmd/excel"
	"github.com/YangSen-qn/Kodo/core/log"
)

const (
	resultSheetTitleWidth    = 20
	resultSheetTitleHeight   = 20
	resultSheetContentWidth  = 20
	resultSheetContentHeight = 20
)

var (
	resultSheetTitleStyle   *excel.CellStyle
	resultSheetContentStyle *excel.CellStyle
)

func saveResultItemsToLocalAsExcel(fileName string, items []*log.QueryResultItem) error {

	if items == nil || len(items) == 0 {
		return nil
	}

	sheet := excel.NewSheet(fileName, "Sheet1")
	configResultSheet(sheet)

	row := 1
	titleList := []string{"IP", "国家", "区域", "城市", "数量"}
	for index, title := range titleList {
		titleCell := &excel.Cell{
			Row:    row,
			Column: index,
			Value:  title,
			Style:  resultSheetTitleStyle,
			Width:  resultSheetTitleWidth,
			Height: resultSheetTitleHeight,
		}
		_ = sheet.SetCell(titleCell)
	}

	for h, item := range items {
		if item == nil {
			continue
		}
		itemInfo := []interface{}{item.IP, item.Country, item.Region, item.City, item.Count}

		for v, title := range itemInfo {
			titleCell := &excel.Cell{
				Row:    h + row + 1,
				Column: v,
				Value:  title,
				Style:  resultSheetContentStyle,
				Width:  resultSheetContentWidth,
				Height: resultSheetContentHeight,
			}
			_ = sheet.SetCell(titleCell)
		}
	}

	return sheet.SaveAs(fileName)
}

func configResultSheet(sheet *excel.Sheet) {

	resultSheetTitleStyle = excel.NewCellStyle()
	resultSheetTitleStyle.SetCellStyle(
		excel.FontStyle(excel.BoldOption(true),
			excel.ColorOption("#777777"),
			excel.SizeOption(14), ),
		excel.AlignmentStyle(excel.HorizontalOption(excel.StringCenter),
			excel.VerticalOption(excel.StringCenter)),
		excel.FillStyle(excel.TypeOption(excel.StringPattern),
			excel.ColorsOption("#FFFF88"),
			excel.PatternOption(1)))
	_ = sheet.AddCellStyle(resultSheetTitleStyle)

	resultSheetContentStyle = excel.NewCellStyle()
	resultSheetContentStyle.SetCellStyle(
		excel.FontStyle(excel.ColorOption("#777777"),
			excel.SizeOption(13), ),
		excel.AlignmentStyle(excel.HorizontalOption(excel.StringCenter),
			excel.VerticalOption(excel.StringCenter)))
	_ = sheet.AddCellStyle(resultSheetContentStyle)
}
