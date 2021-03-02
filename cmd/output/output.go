package output

import "fmt"

var globalOutput IOutput = &DefaultOutput{IsColorful: true}

func SetOutput(output IOutput) {
	globalOutput = output
}

// --- output message
type OutputMessageLevel int8

const (
	OutputMessageLevelInfo OutputMessageLevel = iota
	OutputMessageLevelDebug
	OutputMessageLevelWarning
	OutputMessageLevelError
)

type OutputMessageStyle struct {
	width int
	color int // default PrintForegroundColorWhite
}

func NewOutputMessageStyle() OutputMessageStyle {
	return OutputMessageStyle{
		width: 0,
		color: PrintForegroundColorWhite,
	}
}

// 0: 为不限制
func (style OutputMessageStyle) Width(w int) OutputMessageStyle {
	style.width = w
	return style
}

// default PrintForegroundColorWhite
func (style OutputMessageStyle) Color(c int) OutputMessageStyle {
	style.color = c
	return style
}

// --- output
type IOutput interface {
	Output(messageLevel OutputMessageLevel, messageStyle OutputMessageStyle, message interface{})
}

type OutputInfo struct {
	level OutputMessageLevel
}

func I() *OutputInfo {
	return &OutputInfo{
		level: OutputMessageLevelInfo,
	}
}

func D() *OutputInfo {
	return &OutputInfo{
		level: OutputMessageLevelDebug,
	}
}

func W() *OutputInfo {
	return &OutputInfo{
		level: OutputMessageLevelWarning,
	}
}

func E(err error) {
	o := &OutputInfo{
		level: OutputMessageLevelWarning,
	}
	o.Output(ErrorMessage(err))
}

func (o *OutputInfo) Output(message interface{}) {
	o.OutputWithStyle(NewOutputMessageStyle(), message)
}

func (o *OutputInfo) OutputWithStyle(messageStyle OutputMessageStyle, message interface{}) {
	if message != nil {
		globalOutput.Output(OutputMessageLevelInfo, messageStyle, message)
	}
}

func (o *OutputInfo) OutputFormat(format string, args ...interface{}) {
	o.OutputFormatWithStyle(NewOutputMessageStyle(), format, args...)
}

func (o *OutputInfo) OutputFormatWithStyle(messageStyle OutputMessageStyle, format string, args ...interface{}) {
	o.OutputWithStyle(messageStyle, fmt.Sprintf(format, args...))
}

// --- error output message
type errorMessage struct {
	Error string `json:"error"`
}

func (e *errorMessage) String() string {
	return e.Error
}

func ErrorMessage(err error) interface{} {
	return &errorMessage{Error: err.Error()}
}
