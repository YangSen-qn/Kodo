package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

// Sheet Cell
type Cell struct {
	Row    int
	Column int
	axis   string
	Value  interface{}
	Style  *CellStyle
	Width  float64
	Height float64
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
		fmt.Println("open error:", err)
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
		sheet.file.SetRowHeight(sheet.name, cell.Row, cell.Height)
	}

	if err == nil && cell.Style.id > 0 {
		err = sheet.file.SetCellStyle(sheet.name, axis, axis, cell.Style.id)
	}

	return err
}

func (sheet *Sheet) AddCellStyle(style *CellStyle) error {
	id, err := sheet.file.NewStyle(style.style)
	if err != nil {
		return err
	}
	style.id = id
	return nil
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
