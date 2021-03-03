package uplog

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/YangSen-qn/Kodo/cmd/excel"
	"github.com/YangSen-qn/Kodo/core/log"
)

const (
	defaultCellHeight = 20
	emptyCellWidth    = 4
	titleCellWidth    = 30
	nameCellWidth     = 33
	countCellWidth    = 12
	percentCellWidth  = 12

	excelCellTypeEmpty   = 0
	excelCellTypeTitle   = 1
	excelCellTypeName    = 2
	excelCellTypeCount   = 3
	excelCellTypePercent = 4
)

var (
	emptyCellStyleId   = -1
	titleCellStyleId   = -1
	nameCellStyleId    = -1
	countCellStyleId   = -1
	percentCellStyleId = -1
)

func saveVersionToLocalAsExcel(fileName string, sdkName string, allVersionLogCount int, versionList [] *log.QueryResultVersion, types []string) {

	if versionList == nil || len(versionList) == 0 {
		return
	}

	sheetName := sdkName

	// 删除无用的Sheet1
	excel.DeleteSheet(fileName, "Sheet1")
	excel.DeleteSheet(fileName, sheetName)

	sheet := excel.NewSheet(fileName, sheetName)
	configVersionSheet(sheet)

	for i, sdkVersionInfo := range versionList {

		sdkVersion := sdkVersionInfo.Version()

		row := 1
		columnStart := i*3 + i

		// 头
		column := columnStart + 1
		totalPercent := log.CalculatePercent(sdkVersionInfo.TotalCount(), allVersionLogCount)
		versionDescription := sdkName + " " + sdkVersion + "(" + percentToString(totalPercent) + ")"
		_ = sheet.SetCell(newCell(excelCellTypeEmpty, row, columnStart, ""))
		_ = sheet.SetCell(newCell(excelCellTypeTitle, row, column, versionDescription))
		_ = sheet.MergeCell(column, row, column+2, row)

		// total success dns
		nameList := []string{sdkVersion + " totalCount", sdkVersion + " successCount", sdkVersion + " dnsErrorCount"}
		countList := []interface{}{sdkVersionInfo.TotalCount(), sdkVersionInfo.SuccessCount(), sdkVersionInfo.DnsErrorCount()}
		percentList := []float64{-1, sdkVersionInfo.SuccessPercent(), sdkVersionInfo.DnsErrorPercent()}
		for j := 0; j < len(nameList); j++ {
			row++
			_ = sheet.SetCell(newCell(excelCellTypeEmpty, row, columnStart, ""))
			_ = sheet.SetCell(newCell(excelCellTypeName, row, columnStart+1, nameList[j]))
			_ = sheet.SetCell(newCell(excelCellTypeCount, row, columnStart+2, countList[j]))
			if percentList[j] >= 0 {
				_ = sheet.SetCell(newCell(excelCellTypePercent, row, columnStart+3, percentToString(percentList[j])))
			}
		}

		// other type
		for _, key := range types {
			value := sdkVersionInfo.TypeInfo(key)
			percent := sdkVersionInfo.TypeInfoPercent(value)

			row++
			column = columnStart
			_ = sheet.SetCell(newCell(excelCellTypeEmpty, row, columnStart, ""))
			_ = sheet.SetCell(newCell(excelCellTypeName, row, columnStart+1, sdkVersion+" "+key))
			_ = sheet.SetCell(newCell(excelCellTypeCount, row, columnStart+2, value.TotalCount()))
			_ = sheet.SetCell(newCell(excelCellTypePercent, row, columnStart+3, percentToString(percent)))
		}
	}

	err := sheet.SaveAs(fileName);
	if err != nil {
	}
}

func configVersionSheet(sheet *excel.Sheet) {

	borderColor := "#DDDDDD"
	defaultBorder := excel.Border(borderColor, borderColor, borderColor, borderColor)
	titleFill := excel.Fill("#FFFF88")
	titleFont := excel.BoldFont(14, "")
	contentFont := excel.Font(13, "")
	leftAlignment := excel.Alignment(excel.AlignmentHorizontalLeft, excel.AlignmentCenter)
	centerAlignment := excel.Alignment(excel.AlignmentCenter, excel.AlignmentCenter)
	rightAlignment := excel.Alignment(excel.AlignmentHorizontalRight, excel.AlignmentCenter)

	titleCellStyle := &excelize.Style{
		Border:        defaultBorder,
		Fill:          titleFill,
		Font:          titleFont,
		Alignment:     centerAlignment,
	}
	titleCellStyleId, _ = sheet.AddCellStyle(titleCellStyle)

	nameCellStyle := &excelize.Style{
		Border:        defaultBorder,
		Font:          contentFont,
		Alignment:     leftAlignment,
	}
	nameCellStyleId, _ = sheet.AddCellStyle(nameCellStyle)

	countCellStyle := &excelize.Style{
		Border:        defaultBorder,
		Font:          contentFont,
		Alignment:     rightAlignment,
	}
	countCellStyleId, _ = sheet.AddCellStyle(countCellStyle)
	percentCellStyleId = countCellStyleId
}

func newCell(typeId int, row int, column int, value interface{}) *excel.Cell {

	var styleId = -1

	cellWidth := 0.0
	cellHeight := 0.0
	if typeId == excelCellTypeEmpty {
		cellWidth = emptyCellWidth
		cellHeight = defaultCellHeight
	} else if typeId == excelCellTypeTitle {
		styleId = titleCellStyleId
		cellWidth = titleCellWidth
		cellHeight = defaultCellHeight
	} else if typeId == excelCellTypeName {
		styleId = nameCellStyleId
		cellWidth = nameCellWidth
		cellHeight = defaultCellHeight
	} else if typeId == excelCellTypeCount {
		styleId = countCellStyleId
		cellWidth = countCellWidth
		cellHeight = defaultCellHeight
	} else {
		styleId = percentCellStyleId
		cellWidth = percentCellWidth
		cellHeight = defaultCellHeight
	}

	return &excel.Cell{
		Row:     row,
		Column:  column,
		Value:   value,
		StyleId: styleId,
		Width:   cellWidth,
		Height:  cellHeight,
	}
}

func percentToString(percent float64) string {
	return fmt.Sprintf("%.4f%%", percent*100)
}
