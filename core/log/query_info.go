package log

import "fmt"

func QueryInfoOfCannotConnectToServer(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	param.QueryString = QueryTypeQueryString(param.SDKVersion,
		param.SDKType,
		[]string{QS_ResultTimeout, QS_ResultCannotConnectToHost, QS_ResultUnknownHost, QS_ResultNetworkError})
	fmt.Println("start time:", param.StartTime)
	fmt.Println("end   time:", param.EndTime)
	fmt.Println("query string:", param.QueryString)
	QueryInfo(param, partResultChan, errorResultChan)
}

func QueryInfo(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	if len(param.QueryString) == 0 {
		return
	}
	param.check()
	queryInfoSeparateByPage(param, partResultChan, errorResultChan)
}
