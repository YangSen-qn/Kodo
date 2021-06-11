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

type SuccessRateXlsx struct {
	filePath   string
	sheetName  string
	currentRow int
	sheet      *excel.Sheet
}

func NewSuccessRateXlsx(filePath, sheet string) *SuccessRateXlsx {
	s := &SuccessRateXlsx{
		filePath:   filePath,
		sheetName:  sheet,
		currentRow: 0,
	}
	s.Setup()
	return s
}

func (s *SuccessRateXlsx) Setup() {
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
		Value:  "总量",
		Width:  20,
		Height: 15,
	})

	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 2,
		Value:  "成功量",
		Width:  20,
		Height: 15,
	})

	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 3,
		Value:  "服务异常量",
		Width:  20,
		Height: 15,
	})

	s.save()
}

func (s *SuccessRateXlsx) Write(info log.QuerySuccessRateInfo) error {
	if err := s.createFileIfNotExist(); err != nil {
		return err
	}

	start := util.GetDateStringWithTimestamp(info.Start)
	end := util.GetDateStringWithTimestamp(info.End)
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
		Value:  info.AllCount,
		Width:  10,
		Height: 15,
	})

	successRow := fmt.Sprintf("%d(%.4f%%)", info.SuccessCount, float64(info.SuccessCount)/float64(info.SuccessCount))
	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 2,
		Value:  successRow,
		Width:  10,
		Height: 15,
	})

	serverErrorRow := fmt.Sprintf("%d(%.4f%%)", info.ServerErrorCount, float64(info.ServerErrorCount)/float64(info.SuccessCount))
	s.sheet.SetCell(&excel.Cell{
		Row:    s.currentRow,
		Column: 3,
		Value:  serverErrorRow,
		Width:  10,
		Height: 15,
	})

	return s.save()
}

func (s *SuccessRateXlsx) Complete() error {
	return s.save()
}

func (s *SuccessRateXlsx) save() error {
	return s.sheet.SaveAs(s.filePath)
}

func (s *SuccessRateXlsx) createFileIfNotExist() error {
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
