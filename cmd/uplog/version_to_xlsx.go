package uplog

import (
	"fmt"
	"github.com/YangSen-qn/Kodo/core/log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
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
	emptyCellStyleId   int = -1
	titleCellStyleId   int = -1
	nameCellStyleId    int = -1
	countCellStyleId   int = -1
	percentCellStyleId int = -1
)

type excelCell struct {
	sheet      string
	typeId     int
	row        int
	column     string
	value      interface{}
	styleId    int
	cellWidth  float64
	cellHeight float64
}

func saveVersionToLocalAsExcel(fileName string, sdkName string, allVersionLogCount int, versionList [] *log.QueryResultVersion, types []string) {

	if versionList == nil || len(versionList) == 0 {
		return
	}

	sheetName := sdkName
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		file = excelize.NewFile()
	}
	configFile(file)

	// 删除无用的Sheet1
	if file.GetSheetIndex("Sheet1") > -1 {
		file.DeleteSheet("Sheet1")
	}

	if file.GetSheetIndex(sheetName) > 0 {
		file.DeleteSheet(sheetName)
	}
	file.NewSheet(sheetName)

	for i, sdkVersionInfo := range versionList {

		sdkVersion := sdkVersionInfo.Version()

		// 头
		row := 1
		column := i*3 + i
		totalPercent := log.CalculatePercent(sdkVersionInfo.TotalCount(), allVersionLogCount)
		tileCell0 := newCell(excelCellTypeEmpty, sheetName, row, getExcelCellColumnString(column), "");
		column++
		tileCell1 := newCell(excelCellTypeTitle, sheetName, row, getExcelCellColumnString(column), sdkName+" "+sdkVersion+"("+percentToString(totalPercent)+")");
		column++
		tileCell2 := newCell(excelCellTypeTitle, sheetName, row, getExcelCellColumnString(column), "");
		column++
		tileCell3 := newCell(excelCellTypeTitle, sheetName, row, getExcelCellColumnString(column), "");
		column++
		setCellValue(file, tileCell0)
		setCellValue(file, tileCell1)
		setCellValue(file, tileCell2)
		setCellValue(file, tileCell3)
		err = file.MergeCell(sheetName, excelCellAxis(tileCell1), excelCellAxis(tileCell3))
		if err != nil {
		}

		// 总量
		row += 1
		column = i*3 + i
		setCellValue(file, newCell(excelCellTypeEmpty, sheetName, row, getExcelCellColumnString(column), ""));
		column++
		setCellValue(file, newCell(excelCellTypeName, sheetName, row, getExcelCellColumnString(column), sdkVersion+" totalCount"));
		column++
		setCellValue(file, newCell(excelCellTypeCount, sheetName, row, getExcelCellColumnString(column), sdkVersionInfo.TotalCount()));
		column++

		// 成功数据
		row += 1
		column = i*3 + i
		setCellValue(file, newCell(excelCellTypeEmpty, sheetName, row, getExcelCellColumnString(column), ""));
		column++
		setCellValue(file, newCell(excelCellTypeName, sheetName, row, getExcelCellColumnString(column), sdkVersion+" successCount"));
		column++
		setCellValue(file, newCell(excelCellTypeCount, sheetName, row, getExcelCellColumnString(column), sdkVersionInfo.SuccessCount()));
		column++
		setCellValue(file, newCell(excelCellTypePercent, sheetName, row, getExcelCellColumnString(column), percentToString(sdkVersionInfo.SuccessPercent())));
		column++

		// dns 数据
		row += 1
		column = i*3 + i
		setCellValue(file, newCell(excelCellTypeEmpty, sheetName, row, getExcelCellColumnString(column), ""));
		column++
		setCellValue(file, newCell(excelCellTypeName, sheetName, row, getExcelCellColumnString(column), sdkVersion+" dnsErrorCount"));
		column++
		setCellValue(file, newCell(excelCellTypeCount, sheetName, row, getExcelCellColumnString(column), sdkVersionInfo.DnsErrorCount()));
		column++
		setCellValue(file, newCell(excelCellTypePercent, sheetName, row, getExcelCellColumnString(column), percentToString(sdkVersionInfo.DnsErrorPercent())));
		column++

		for _, key := range types {
			value := sdkVersionInfo.TypeInfo(key)
			percent := sdkVersionInfo.TypeInfoPercent(value)

			row += 1
			column = i*3 + i
			setCellValue(file, newCell(excelCellTypeEmpty, sheetName, row, getExcelCellColumnString(column), ""));
			column++
			setCellValue(file, newCell(excelCellTypeName, sheetName, row, getExcelCellColumnString(column), sdkVersion+" "+key));
			column++
			setCellValue(file, newCell(excelCellTypeCount, sheetName, row, getExcelCellColumnString(column), value.TotalCount()));
			column++
			setCellValue(file, newCell(excelCellTypePercent, sheetName, row, getExcelCellColumnString(column), percentToString(percent)));
			column++
		}
	}

	err = file.SaveAs(fileName);
	if err != nil {
	}
}

func configFile(file *excelize.File) {
	fontStyle := `"font":{"bold":false,"italic":true,"family":"Times","size":14,"color":"#777777"}`
	boldFontStyle := `"font":{"bold":true,"italic":true,"family":"Times","size":14,"color":"#777777"}`
	leftAlignmentStyle := `"alignment":{"horizontal":"left","Vertical":"center"}`
	centerAlignmentStyle := `"alignment":{"horizontal":"center","Vertical":"center"}`
	rightAlignmentStyle := `"alignment":{"horizontal":"right","Vertical":"center"}`
	borderStyle := `"border":[{"type":"left","color":"FF0000","style":1}, {"type":"top","color":"FF0000","style":1}, {"type":"right","color":"FF0000","style":1}, {"type":"bottom","color":"FF0000","style":1}]`
	yellowFillStyle := `"fill":{"type":"pattern","color":["#FFFF88"],"pattern":1}`

	var err error
	titleCellStyleId, err = file.NewStyle("{" + boldFontStyle + "," + centerAlignmentStyle + "," + borderStyle + "," + yellowFillStyle + "}")
	nameCellStyleId, err = file.NewStyle("{" + fontStyle + "," + leftAlignmentStyle + "}")
	countCellStyleId, err = file.NewStyle("{" + fontStyle + "," + rightAlignmentStyle + "}")
	percentCellStyleId, err = file.NewStyle("{" + fontStyle + "," + rightAlignmentStyle + "}")

	if err != nil {
	}
}

func newCell(typeId int, sheet string, row int, column string, value interface{}) *excelCell {

	styleId := -1
	cellWidth := 0.0
	cellHeight := 0.0
	if typeId == excelCellTypeEmpty {
		styleId = emptyCellStyleId
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

	return &excelCell{
		sheet:      sheet,
		typeId:     typeId,
		row:        row,
		column:     column,
		value:      value,
		styleId:    styleId,
		cellWidth:  cellWidth,
		cellHeight: cellHeight,
	}
}

func excelCellAxis(cell *excelCell) string {
	return cell.column + strconv.Itoa(cell.row)
}

func setCellValue(file *excelize.File, cell *excelCell) {
	var err error

	axis := excelCellAxis(cell)
	if stringValue, ok := cell.value.(string); ok {
		err = file.SetCellValue(cell.sheet, axis, stringValue)
	} else if intValue, ok := cell.value.(int); ok {
		err = file.SetCellInt(cell.sheet, axis, intValue)
	} else if float64Value, ok := cell.value.(float64); ok {
		err = file.SetCellFloat(cell.sheet, axis, float64Value, 4, 32)
	} else if float32Value, ok := cell.value.(float32); ok {
		err = file.SetCellFloat(cell.sheet, axis, float64(float32Value), 4, 32)
	}

	if err == nil {
		err = file.SetRowHeight(cell.sheet, cell.row, cell.cellHeight)
	}

	if err == nil {
		err = file.SetColWidth(cell.sheet, cell.column, cell.column, cell.cellWidth)
	}

	if err == nil && cell.styleId > 0 {
		err = file.SetCellStyle(cell.sheet, axis, axis, cell.styleId)
	}

	if err != nil {
	}
}

func getExcelCellColumnString(column int) string {

	cellColumnBase := []string{
		"A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X",
		"Y", "Z"}

	columnString := cellColumnBase[column%len(cellColumnBase)]
	for column >= len(cellColumnBase) {
		column = column/len(cellColumnBase) - 1
		columnString = cellColumnBase[column%len(cellColumnBase)] + columnString
	}

	return columnString
}

func percentToString(percent float64) string {
	return fmt.Sprintf("%.4f%%", percent*100)
}

// 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func isFile(path string) bool {
	return !isDir(path)
}
