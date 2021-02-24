package output

import (
	"fmt"
	"strings"
)

const (
	PrintShowStyleDefault   = 0
	PrintShowStyleHighlight = 1
	PrintShowStyleUnderline = 4
	PrintShowStyleTwinkle   = 5
	PrintShowStyleInverse   = 7
	PrintShowStyleInvisible = 8

	PrintForegroundColorBlack       = 30
	PrintForegroundColorRed         = 31
	PrintForegroundColorGreen       = 32
	PrintForegroundColorYellow      = 33
	PrintForegroundColorBlue        = 34
	PrintForegroundColorPurple      = 35
	PrintForegroundColorUltramarine = 36
	PrintForegroundColorWhite       = 37

	PrintBackgroundColorNone        = 1
	PrintBackgroundColorBlack       = 40
	PrintBackgroundColorRed         = 41
	PrintBackgroundColorGreen       = 42
	PrintBackgroundColorYellow      = 43
	PrintBackgroundColorBlue        = 44
	PrintBackgroundColorPurple      = 45
	PrintBackgroundColorUltramarine = 46
	PrintBackgroundColorWhite       = 47

	// 显示方式，前景色，背景色，输出字符串
	colorfulPrintColorFormat = "\033[%d;%d;%dm%s\033[0m"
	defaultPrintColorFormat  = "%s"
)

type PrintFormat struct {
	IsColorful      bool
	Style           int
	ForegroundColor int
	BackgroundColor int
	width           int
}

func NewPrintFormat() *PrintFormat {
	return &PrintFormat{
		IsColorful:      true,
		Style:           PrintShowStyleDefault,
		ForegroundColor: PrintForegroundColorWhite,
		BackgroundColor: PrintBackgroundColorNone,
		width:           0,
	}
}

func (format *PrintFormat) check(s string) {
	if format.width < 1 {
		format.width = len(s)
	}

	if format.Style < 0 || format.Style > 8 ||
		format.Style == 2 || format.Style == 3 || format.Style == 6 {
		format.Style = PrintShowStyleDefault
	}

	if format.ForegroundColor < 30 || format.ForegroundColor > 37 {
		format.ForegroundColor = PrintForegroundColorWhite
	}

	if format.BackgroundColor < 40 || format.BackgroundColor > 47 {
		format.BackgroundColor = PrintBackgroundColorNone
	}
}

func printBeautiful(s string, format *PrintFormat) {

	format.check(s)

	stringLen := len(s)
	if stringLen < format.width {
		s = s + strings.Repeat(" ", format.width-stringLen)
	} else if stringLen > format.width {
		s = s[0:format.width]
	}
	if format.IsColorful {
		fmt.Printf(colorfulPrintColorFormat, format.Style, format.BackgroundColor, format.ForegroundColor, s)
	} else {
		fmt.Printf(defaultPrintColorFormat, s)
	}

}

func printBeautifulNewLine() {
	fmt.Println("")
}