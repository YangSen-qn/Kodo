package output

import "fmt"

type DefaultOutput struct {
	IsColorful bool
}

func (output *DefaultOutput) Output(messageLevel OutputMessageLevel, messageStyle OutputMessageStyle, message interface{}) {

	format := NewPrintFormat()
	format.IsColorful = output.IsColorful
	format.width = messageStyle.width
	msg := fmt.Sprintf("%s", message)
	if messageLevel == OutputMessageLevelWarning {
		format.ForegroundColor = PrintForegroundColorYellow
	} else if messageLevel == OutputMessageLevelError {
		format.ForegroundColor = PrintForegroundColorRed
	} else {
		format.ForegroundColor = messageStyle.color
	}
	printBeautiful(msg, format)
}
