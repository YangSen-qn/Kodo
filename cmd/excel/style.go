package excel

import (
	"fmt"
	"strings"
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

type CellStyle struct {
	id    int
	style string
}

func NewCellStyle() *CellStyle {
	return &CellStyle{id: -1}
}

type Style string
type Option string

func joinOptions(options ...Option) string {
	optionsReal := make([]string, 0, len(options))
	for _, option := range options {
		if len(option) > 0 {
			optionsReal = append(optionsReal, string(option))
		}
	}
	return strings.Join(optionsReal, ",")
}

// font style
func BoldOption(enable bool) Option {
	if enable {
		return `"bold":true`
	} else {
		return `"bold":false`
	}
}

func ItalicOption(enable bool) Option {
	if enable {
		return `"italic":true`
	} else {
		return `"italic":false`
	}
}

const (
	StringTimes Style = "Times"
)

func FamilyOption(family Style) Option {
	return Option(fmt.Sprintf(`"family":"%s"`, family))
}

func SizeOption(size int) Option {
	return Option(fmt.Sprintf(`"size":%d`, size))
}

func ColorOption(color string) Option {
	return Option(fmt.Sprintf(`"color":"%s"`, color))
}

func FontStyle(options ...Option) Style {
	styleString := joinOptions(options...)
	if len(styleString) == 0 {
		return ""
	} else {
		return Style(`"font":{` + styleString + `}`)
	}
}

// alignment style
const (
	StringLeft   Option = "left"
	StringRight  Option = "right"
	StringCenter Option = "center"
	StringTop    Option = "top"
	StringBottom Option = "bottom"
)

func HorizontalOption(alignment Option) Option {
	return Option(fmt.Sprintf(`"horizontal":"%s"`, alignment))
}

func VerticalOption(alignment Option) Option {
	return Option(fmt.Sprintf(`"vertical":"%s"`, alignment))
}

func AlignmentStyle(options ...Option) Style {
	optionString := joinOptions(options...)
	if len(optionString) == 0 {
		return ""
	} else {
		return Style(`"alignment":{` + optionString + `}`)
	}
}

// fill style
func PatternOption(pattern int) Option {
	return Option(fmt.Sprintf(`"pattern":%d`, pattern))
}

const (
	StringPattern Option = "pattern"
)

func TypeOption(typeOption Option) Option {
	return Option(fmt.Sprintf(`"type":"%s"`, typeOption))
}

func FillStyle(options ...Option) Style {
	optionString := joinOptions(options...)
	if len(optionString) == 0 {
		return ""
	} else {
		return Style(`"fill":{` + optionString + `}`)
	}
}

// border style
func ColorsOption(colors ...string) Option {
	colorOptions := make([]Option, len(colors))
	for _, color := range colors {
		color = `"` + color + `"`
		colorOptions = append(colorOptions, Option(color))
	}
	return Option(fmt.Sprintf(`"color":[%s]`, joinOptions(colorOptions...)))
}

func StyleOption(style int) Option {
	return Option(fmt.Sprintf(`"style":"%d"`, style))
}

func BorderLineOption(options ...Option) Option {
	optionString := joinOptions(options...)
	return Option(fmt.Sprintf(`{%s}`, optionString))
}

func BorderStyle(options ...Option) Style {
	optionString := joinOptions(options...)
	if len(optionString) == 0 {
		return ""
	} else {
		return Style(`"border":[` + optionString + `]`)
	}
}

func (style *CellStyle) SetCellStyle(styles ...Style) {
	stylesReal := make([]string, 0, len(styles))
	for _, style := range styles {
		if len(style) > 0 {
			stylesReal = append(stylesReal, string(style))
		}
	}
	style.style = "{" + strings.Join(stylesReal, ",") + "}"
}
