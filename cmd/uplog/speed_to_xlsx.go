package uplog

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/YangSen-qn/Kodo/cmd/excel"
	"github.com/YangSen-qn/Kodo/core/log"
	"github.com/YangSen-qn/Kodo/core/util"
	"github.com/coreos/etcd/pkg/fileutil"
)

type SpeedXlsx struct {
	filePath   string
	sheetName  string
	currentRow int
	sheet      *excel.Sheet
}

func NewSpeedXlsx(filePath, sheet string) *SpeedXlsx {
	s := &SpeedXlsx{
		filePath:   filePath,
		sheetName:  sheet,
		currentRow: 0,
	}
	s.Setup()
	return s
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

func (s *SpeedXlsx) Write(speedInfo log.QuerySpeedInfo) error {
	if err := s.createFileIfNotExist(); err != nil {
		return err
	}

	start := util.GetDateStringWithTimestamp(speedInfo.Start)
	end := util.GetDateStringWithTimestamp(speedInfo.End)
	time := fmt.Sprintf("%v ~ %v", start, end)

	s.currentRow += 1
	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 0,
		Value:  time,
		Width:  40,
		Height: 15,
	})

	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 1,
		Value:  speedInfo.Speed,
		Width:  10,
		Height: 15,
	})
	return s.save()
}

func (s *SpeedXlsx) Compete() error {
	return s.save()
}

func (s *SpeedXlsx) save() error {
	return s.sheet.SaveAs(s.filePath)
}

func (s *SpeedXlsx) createFileIfNotExist() error {
	if len(s.filePath) == 0 {
		return errors.New("file path is invalid")
	}

	dir := filepath.Dir(s.filePath)
	if !fileutil.Exist(dir) {
		if err := fileutil.CreateDirAll(dir); err != nil {
			return err
		}
	}

	return nil
}
