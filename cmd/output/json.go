package output

import (
	"encoding/json"
)

type JsonOutput struct {
	Pretty bool
}

func (o *JsonOutput) Output(messageLevel OutputMessageLevel, messageStyle OutputMessageStyle, message interface{}) {

	msg := ""
	msgByte, err := json.Marshal(message)
	if err == nil {
		msg = string(msgByte)
	}

	format := NewPrintFormat()
	format.IsColorful = false
	printBeautiful(msg+"\n", format)
}

//func JsonToMap(jsonStr string) (map[string]interface{}, error) {
//	m := make(map[string]interface{})
//	err := json.Unmarshal([]byte(jsonStr), &m)
//	if err != nil {
//		fmt.Printf("Unmarshal with error: %+v\n", err)
//		return nil, err
//	}
//
//	for k, v := range m {
//		fmt.Printf("%v: %v\n", k, v)
//	}
//
//	return m, nil
//}