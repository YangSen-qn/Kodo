package log

func QueryInfoOfCannotConnectToServer(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	param.QueryString = QueryTypeRequestQueryString(param.SDKVersion,
		param.SDKType,
		[]string{QS_Timeout, QS_CannotConnectToHost, QS_UnknownHost, QS_NetworkError})
	QueryInfo(param, partResultChan, errorResultChan)
}

func QueryInfo(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	if len(param.QueryString) == 0 {
		return
	}
	param.check()
	queryInfoSeparateByPage(param, partResultChan, errorResultChan)
}
