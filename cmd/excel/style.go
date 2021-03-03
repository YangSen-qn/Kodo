package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Excel Cell style
/*
fontStyle := `"font":{"bold":false,"italic":true,"family":"Times","size":14,"color":"#777777"}`
	boldFontStyle := `"font":{"bold":true,"italic":true,"family":"Times","size":14,"color":"#777777"}`
	leftAlignmentStyle := `"alignment":{"horizontal":"left","Vertical":"center"}`
	centerAlignmentStyle := `"alignment":{"horizontal":"center","Vertical":"center"}`
	rightAlignmentStyle := `"alignment":{"horizontal":"right","Vertical":"center"}`
	borderStyle := `"border":[{"type":"left","color":"FF0000","style":1}, {"type":"top","color":"FF0000","style":1}, {"type":"right","color":"FF0000","style":1}, {"type":"bottom","color":"FF0000","style":1}]`
	yellowFillStyle := `"fill":{"type":"pattern","color":["#FFFF88"],"pattern":1}`
*/

// border:
const (
	BorderTypeLeft         = "left"
	BorderTypeRight        = "right"
	BorderTypeTop          = "top"
	BorderTypeBottom       = "bottom"
	BorderTypeDiagonalDown = "diagonalDown"
	BorderTypeDiagonalUp   = "diagonalUp"

	BorderStyle1 = 1 // 实细线
	BorderStyle2 = 2 // 实粗线
	BorderStyle3 = 3 // 半粗半细曲线
	BorderStyle4 = 4 // 均细曲线
	BorderStyle5 = 5 // 实粗线 比2粗
	BorderStyle6 = 6 // 双实线
	BorderStyle7 = 7 // 虚粗线
)

func Border(topColor, leftColor, bottomColor, rightColor string) []excelize.Border {
	return []excelize.Border{
		{Type: BorderTypeLeft, Color: leftColor, Style: BorderStyle1},
		{Type: BorderTypeRight, Color: rightColor, Style: BorderStyle1},
		{Type: BorderTypeTop, Color: topColor, Style: BorderStyle1},
		{Type: BorderTypeBottom, Color: bottomColor, Style: BorderStyle1},
	}
}

// font:
const (
	FontFamilyTimeTimesNewRoman = "Times New Roman"
)

func Font(size float64, color string) *excelize.Font {
	if len(color) == 0 {
		color = "#777777"
	}
	return &excelize.Font{
		Bold:      false,
		Italic:    false,
		Underline: "",
		Family:    FontFamilyTimeTimesNewRoman,
		Size:      size,
		Color:     color,
	}
}

func BoldFont(size float64, color string) *excelize.Font {
	if len(color) == 0 {
		color = "#777777"
	}
	return &excelize.Font{
		Bold:      true,
		Italic:    false,
		Underline: "",
		Family:    FontFamilyTimeTimesNewRoman,
		Size:      size,
		Color:     color,
	}
}

// fill:
const (
	FillTypeGradient = "gradient"
	FillTypePattern  = "pattern"

	FillPattern1 = 1

	FillShading0 = 0
)

func Fill(color string) excelize.Fill {
	return excelize.Fill{
		Type:    FillTypePattern,
		Pattern: FillPattern1,
		Color:   []string{color},
		Shading: FillShading0,
	}
}

// alignment:
const (
	AlignmentCenter          = "center"
	AlignmentHorizontalLeft  = "left"
	AlignmentHorizontalRight = "right"
	AlignmentVerticalTop     = "top"
	AlignmentVerticalBottom  = "bottom"
)

func Alignment(horizontal, vertical string) *excelize.Alignment {
	return &excelize.Alignment{
		Horizontal:      horizontal,
		Indent:          0,
		JustifyLastLine: false,
		ReadingOrder:    0,
		RelativeIndent:  0,
		ShrinkToFit:     false,
		TextRotation:    0,
		Vertical:        vertical,
		WrapText:        false,
	}
}
