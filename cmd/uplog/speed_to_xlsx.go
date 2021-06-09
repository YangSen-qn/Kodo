package uplog

import (
	"fmt"

	"github.com/YangSen-qn/Kodo/cmd/excel"
	"github.com/YangSen-qn/Kodo/core/log"
	"github.com/YangSen-qn/Kodo/core/util"
)

type SpeedXlsx struct {
	filePath   string
	sheetName  string
	currentRow int
	sheet      *excel.Sheet
}

func NewSpeedXlsx(filePath, sheet string) *SpeedXlsx {
	return &SpeedXlsx{
		filePath:   filePath,
		sheetName:  sheet,
		currentRow: 0,
	}
}

func (s *SpeedXlsx) Setup() {
	s.sheet = excel.NewSheetCreateWhileNotExist(s.filePath, s.sheetName)
	s.currentRow += 1
	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 0,
		Value:  "时间",
		Width:  80,
		Height: 15,
	})

	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 1,
		Value:  "速度(KB/S)",
		Width:  20,
		Height: 15,
	})
	s.save()
}

func (s *SpeedXlsx) Write(speedInfo log.QuerySpeedInfo) {
	start := util.GetDateStringWithTimestamp(speedInfo.Start)
	end := util.GetDateStringWithTimestamp(speedInfo.End)
	time := fmt.Sprintf("%v ~ %v", start, end)

	s.currentRow += 1
	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 0,
		Value:  time,
		Width:  80,
		Height: 15,
	})

	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 1,
		Value:  time,
		Width:  20,
		Height: 15,
	})
	s.save()
}

func (s *SpeedXlsx) Compete() {
	s.save()
}

func (s *SpeedXlsx) save() {
	s.sheet.Save()
}
