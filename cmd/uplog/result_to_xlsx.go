package uplog

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
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
	resultSheetTitleStyleId        = -1
	whiteResultSheetContentStyleId = -1
	grayResultSheetContentStyleId  = -1
)

func saveResultItemsToLocalAsExcel(fileName string, items []*log.QueryResultItem) error {

	if items == nil || len(items) == 0 {
		return nil
	}

	sheet := excel.NewSheet(fileName, "Sheet1")
	configResultSheet(sheet)

	row := 1
	titleList := []string{"城市", "区域", "国家", "Host/IpType",
		"ISP", "RemoteIP", "Host", "IP", "数量"}
	for index, title := range titleList {
		titleCell := &excel.Cell{
			Row:     row,
			Column:  index,
			Value:   title,
			StyleId: resultSheetTitleStyleId,
			Width:   resultSheetTitleWidth,
			Height:  resultSheetTitleHeight,
		}
		_ = sheet.SetCell(titleCell)
	}

	lastItemID := ""
	lastItemIDFirstRow := row
	for _, item := range items {
		row++

		if item == nil {
			continue
		}
		id := item.Position()
		remoteNetworkType := item.RemoteNetworkType()
		itemInfo := []interface{}{item.City, item.Region, item.Country, remoteNetworkType, item.ISP,
			item.RemoteIP, item.Host, item.IP, item.Count}

		widthList := []float64{15, 8, 8, 17, 8, 25, 20, 25, 8}
		for v, title := range itemInfo {
			titleCell := &excel.Cell{
				Row:     row,
				Column:  v,
				Value:   title,
				StyleId: whiteResultSheetContentStyleId,
				Width:   widthList[v],
				Height:  resultSheetContentHeight,
			}
			_ = sheet.SetCell(titleCell)
		}

		if lastItemID != id {
			if len(lastItemID) != 0 && row != (lastItemIDFirstRow+1) {
				_ = sheet.MergeCell(0, lastItemIDFirstRow, 0, row-1)
				_ = sheet.MergeCell(1, lastItemIDFirstRow, 1, row-1)
				_ = sheet.MergeCell(2, lastItemIDFirstRow, 2, row-1)
				_ = sheet.SetCellStyle(0, lastItemIDFirstRow, len(titleList)-1, row-1, getNextContentBgStyle())
			} else {
				_ = sheet.SetCellStyle(0, lastItemIDFirstRow, len(titleList)-1, row-1, getNextContentBgStyle())
			}
			lastItemIDFirstRow = row
		}

		lastItemID = id
	}

	return sheet.SaveAs(fileName)
}

func configResultSheet(sheet *excel.Sheet) {

	borderColor := "#DDDDDD"
	defaultBorder := excel.Border(borderColor, borderColor, borderColor, borderColor)
	titleFill := excel.Fill("#FFFF88")
	whiteFill := excel.Fill("#FFFFFF")
	grayFill := excel.Fill("#EEEEEE")
	titleFont  := excel.BoldFont(14, "")
	contentFont  := excel.Font(13, "")
	centerAlignment := excel.Alignment(excel.AlignmentCenter, excel.AlignmentCenter)

	resultSheetTitleStyle := &excelize.Style{
		Border:        defaultBorder,
		Fill:          titleFill,
		Font:          titleFont,
		Alignment:     centerAlignment,
	}
	resultSheetTitleStyleId, _ = sheet.AddCellStyle(resultSheetTitleStyle)

	whiteResultSheetContentStyle := &excelize.Style{
		Border:        defaultBorder,
		Fill:          whiteFill,
		Font:          contentFont,
		Alignment:     centerAlignment,
	}
	whiteResultSheetContentStyleId,_ = sheet.AddCellStyle(whiteResultSheetContentStyle)

	grayResultSheetContentStyle := &excelize.Style{
		Border:        defaultBorder,
		Fill:          grayFill,
		Font:          contentFont,
		Alignment:     centerAlignment,
	}
	grayResultSheetContentStyleId, _ = sheet.AddCellStyle(grayResultSheetContentStyle)
}

var bgStyleSetCount = 0

func getNextContentBgStyle() int {
	if bgStyleSetCount%2 == 0 {
		bgStyleSetCount++
		return whiteResultSheetContentStyleId
	} else {
		bgStyleSetCount++
		return grayResultSheetContentStyleId
	}
}
