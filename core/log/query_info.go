package log

func QueryInfoOfCannotConnectToServer(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	param.QueryString = QueryTypeQueryString(param.SDKVersion,
		param.SDKType,
		[]string{QS_ResultTimeout, QS_ResultCannotConnectToHost, QS_ResultUnknownHost, QS_ResultNetworkError})
	QueryInfo(param, partResultChan, errorResultChan)
}

func QueryInfo(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	if len(param.QueryString) == 0 {
		return
	}
	param.check()
	queryInfoSeparateByPage(param, partResultChan, errorResultChan)
}
