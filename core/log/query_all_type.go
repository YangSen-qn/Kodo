package log

import (
	"fmt"
	"strings"
)

type Range struct {
	Location int `json:"location"`
	Length   int `json:"length"`
}

func (r *Range) Max() int {
	return r.Location + r.Length
}

type TypeParam struct {
	Key        string `json:"key"`
	ValueRange *Range `json:"region"`
}

func (t *TypeParam) handleValue(value interface{}) string {
	if value == nil {
		return ""
	}

	valueString, ok := value.(string)
	if !ok {
		valueString = fmt.Sprintf("%v", t.ValueRange)
	}

	if t.ValueRange != nil && t.ValueRange.Max() < len(valueString) {
		valueString = valueString[t.ValueRange.Location:t.ValueRange.Length]
	}
	return valueString
}

type TypeInfo struct {
	TypeString string
	Count      int
}

func QueryAllType(typeParamList []*TypeParam, param *QueryParam) map[string][]*TypeInfo {
	if param == nil {
		param = &QueryParam{}
	}
	param.check()

	partResultChan := make(chan *QueryResult)
	errorResultChan := make(chan error)
	go QueryInfo(param, partResultChan, errorResultChan)

	allTypeInfo := make(map[string]*TypeInfo)
	for result := range partResultChan {
		for _, typeParam := range typeParamList {
			for _, resultItem := range result.itemList {
				typeValue := resultItem.GetValueByKey(typeParam.Key)
				typeValueString := typeParam.handleValue(typeValue)
				if typeValue != nil {
					key := typeParam.Key + "_" + typeValueString
					info := allTypeInfo[key]
					if info == nil {
						info = &TypeInfo{TypeString: typeValueString}
						allTypeInfo[key] = info
					}
					info.Count += 1
				}
			}
		}
	}

	ret := make(map[string][]*TypeInfo)
	for _, typeParam := range typeParamList {
		infoList := make([]*TypeInfo, 0, 4)
		for key, value := range allTypeInfo {
			if strings.Contains(key, typeParam.Key) {
				infoList = append(infoList, value)
			}
		}
		ret[typeParam.Key] = infoList
	}

	return ret
}
