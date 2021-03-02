package log

func QueryByQueryString(param *QueryParam) (result *QueryResult, err error) {
	if len(param.QueryString) == 0 {
		return
	}
	param.check()
	return queryCountSeparateByDuration(param)
}

func QueryVersion(param *QueryParam, types []string) *QueryResultVersion {

	param.check()

	if types == nil {
		types = QueryType_All()
	}

	version := &QueryResultVersion{
		version: param.SDKVersion,
	}

	allType := types
	allTypeLen := len(allType)
	allTypeChan := []chan int{}
	for i := 0; i < allTypeLen; i++ {
		typeString := allType[i]
		typeChan := make(chan int)
		allTypeChan = append(allTypeChan, typeChan)
		go queryVersionTypeByParam(param, []string{typeString}, typeChan)
	}

	for i := 0; i < allTypeLen; i++ {
		typeString := allType[i]
		typeChan := allTypeChan[i]
		typeCount := <-typeChan

		typeLogInfo := &QueryResultType{
			totalCount: typeCount,
			typeString: typeString,
			itemList:   nil,
		}

		version.addTypeInfo(typeLogInfo)
	}

	version.calculateCount()

	return version
}

func queryVersionTypeByParam(param *QueryParam, typeString []string, count chan int) {
	param.QueryString = QueryTypeQualityQueryString(param.SDKVersion, param.SDKType, typeString)
	result, _ := queryCountSeparateByDuration(param)
	count <- result.totalCount
}
