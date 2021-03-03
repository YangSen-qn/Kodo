package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"strconv"
)

// Sheet Cell
type Cell struct {
	Row     int
	Column  int
	axis    string
	Value   interface{}
	StyleId int
	Width   float64
	Height  float64
}

func (cell *Cell) Axis() string {
	if len(cell.axis) == 0 {
		cell.axis = getExcelCellColumnString(cell.Column) + strconv.Itoa(cell.Row)
	}
	return cell.axis
}

type Sheet struct {
	file *excelize.File
	name string
}

func NewSheet(filePath string, name string) *Sheet {
	if len(name) == 0 {
		name = "Sheet1"
	}

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		file = excelize.NewFile()
	}

	if file.GetSheetIndex(name) == -1 {
		file.NewSheet(name)
	}

	return &Sheet{
		file: file,
		name: name,
	}
}

func DeleteSheet(filePath string, name string) {
	if len(name) == 0 {
		name = "Sheet1"
	}

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return
	}

	if file.GetSheetIndex(name) >= 0 {
		file.DeleteSheet(name)
	}
}

func (sheet *Sheet) SetCell(cell *Cell) error {
	var err error

	axis := cell.Axis()
	if stringValue, ok := cell.Value.(string); ok {
		err = sheet.file.SetCellValue(sheet.name, axis, stringValue)
	} else if intValue, ok := cell.Value.(int); ok {
		err = sheet.file.SetCellInt(sheet.name, axis, intValue)
	} else if float64Value, ok := cell.Value.(float64); ok {
		err = sheet.file.SetCellFloat(sheet.name, axis, float64Value, 4, 32)
	} else if float32Value, ok := cell.Value.(float32); ok {
		err = sheet.file.SetCellFloat(sheet.name, axis, float64(float32Value), 4, 32)
	}

	if err == nil {
		column := getExcelCellColumnString(cell.Column)
		err = sheet.file.SetColWidth(sheet.name, column, column, cell.Width)
	}
	if err == nil {
		err = sheet.file.SetRowHeight(sheet.name, cell.Row, cell.Height)
	}

	if err == nil && cell.StyleId > 0 {
		err = sheet.file.SetCellStyle(sheet.name, axis, axis, cell.StyleId)
	}

	return err
}

func (sheet *Sheet) MergeCell(fromColumn, fromRow, toColumn, toRow int) error {
	fromAxis := getExcelCellColumnString(fromColumn) + strconv.Itoa(fromRow)
	toAxis := getExcelCellColumnString(toColumn) + strconv.Itoa(toRow)
	return sheet.file.MergeCell(sheet.name, fromAxis, toAxis)
}

func (sheet *Sheet) SetCellStyle(fromColumn, fromRow, toColumn, toRow int, styleId int) error {
	fromAxis := getExcelCellColumnString(fromColumn) + strconv.Itoa(fromRow)
	toAxis := getExcelCellColumnString(toColumn) + strconv.Itoa(toRow)
	return sheet.file.SetCellStyle(sheet.name, fromAxis, toAxis, styleId)
}

func (sheet *Sheet) AddCellStyle(style interface{}) (int, error) {
	return sheet.file.NewStyle(style)
}

func (sheet *Sheet) SaveAs(fileName string) error {
	return sheet.file.SaveAs(fileName)
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
