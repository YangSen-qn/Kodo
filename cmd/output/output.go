package output

import (
	"encoding/json"
	"fmt"
)

var globalOutput IOutput = &DefaultOutput{IsColorful:true}

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

type IOutputMessage interface {
	fmt.Stringer
}

// --- string output message
type stringMessage struct {
	Info string `json:"info"`
}

func (s *stringMessage) String() string {
	return s.Info
}

func StringOutputMessage(format string, args ...interface{}) IOutputMessage {
	info := fmt.Sprintf(format, args...)
	return &stringMessage{Info: info}
}

// --- error output message
type errorIOutputMessage struct {
	Error string `json:"error"`
}

func (e errorIOutputMessage) String() string {
	return e.Error
}

func ErrorOutputMessage(err error) IOutputMessage {
	return &errorIOutputMessage{Error: err.Error()}
}

// --- output
type IOutput interface {
	Output(messageLevel OutputMessageLevel, messageStyle OutputMessageStyle, message IOutputMessage)
}

func Info(message IOutputMessage) {
	InfoWithStyle(message, NewOutputMessageStyle())
}

func InfoStringFormat(format string, args ...interface{}) {
	InfoWithStyle(StringOutputMessage(format, args...), NewOutputMessageStyle())
}

func InfoStringFormatWithStyle(messageStyle OutputMessageStyle, format string, args ...interface{}) {
	InfoWithStyle(StringOutputMessage(format, args...), messageStyle)
}

func InfoWithStyle(message IOutputMessage, messageStyle OutputMessageStyle) {
	if message != nil {
		globalOutput.Output(OutputMessageLevelInfo, messageStyle, message)
	}
}

func Debug(message IOutputMessage) {
	DebugWithStyle(message, NewOutputMessageStyle())
}

func DebugStringFormat(format string, args ...interface{}) {
	DebugWithStyle(StringOutputMessage(format, args...), NewOutputMessageStyle())
}

func DebugStringFormatWithStyle(messageStyle OutputMessageStyle, format string, args ...interface{}) {
	DebugWithStyle(StringOutputMessage(format, args...), messageStyle)
}

func DebugWithStyle(message IOutputMessage, messageStyle OutputMessageStyle) {
	if message != nil {
		globalOutput.Output(OutputMessageLevelDebug, messageStyle, message)
	}
}

func Warning(message IOutputMessage) {
	WarningWithStyle(message, NewOutputMessageStyle())
}

func WarningStringFormat(format string, args ...interface{}) {
	WarningWithStyle(StringOutputMessage(format, args...), NewOutputMessageStyle())
}

func WarningStringFormatWithStyle(messageStyle OutputMessageStyle, format string, args ...interface{}) {
	WarningWithStyle(StringOutputMessage(format, args...), messageStyle)
}

func WarningWithStyle(message IOutputMessage, messageStyle OutputMessageStyle) {
	if message != nil {
		globalOutput.Output(OutputMessageLevelWarning, messageStyle, message)
	}
}

func Err(err error) {
	if err != nil {
		globalOutput.Output(OutputMessageLevelError, NewOutputMessageStyle(), ErrorOutputMessage(err))
	}
}

// --- json
type JsonOutput struct {
	pretty bool
}

func (o *JsonOutput) Output(messageLevel OutputMessageLevel, messageStyle OutputMessageStyle, message IOutputMessage) {

	msg := ""
	msgByte, err := json.Marshal(message)
	if err == nil {
		msg = string(msgByte)
	}

	fmt.Println(msg)
	format := NewPrintFormat()
	format.IsColorful = false
	printBeautiful(msg, format)
}

// --- default
type DefaultOutput struct {
	IsColorful bool
}

func (output *DefaultOutput) Output(messageLevel OutputMessageLevel, messageStyle OutputMessageStyle, message IOutputMessage) {

	format := NewPrintFormat()
	format.IsColorful = output.IsColorful
	format.width = messageStyle.width
	msg := message.String()
	if messageLevel == OutputMessageLevelWarning {
		format.ForegroundColor = PrintForegroundColorYellow
	} else if messageLevel == OutputMessageLevelError {
		format.ForegroundColor = PrintForegroundColorRed
	} else {
		format.ForegroundColor = messageStyle.color
	}
	printBeautiful(msg, format)
}
