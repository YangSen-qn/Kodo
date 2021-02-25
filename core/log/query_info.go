package log

import "fmt"

func QueryNetworkSlowInfo(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	param.QueryString = QueryTypeQueryString(param.SDKVersion, param.SDKType, []string{QS_ResultNetworkSlow})
	fmt.Println("query string a:", param.QueryString)
	QueryInfo(param, partResultChan, errorResultChan)
}

func QueryInfo(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	if len(param.QueryString) == 0 {
		return
	}
	param.check()
	queryInfoSeparateByPage(param, partResultChan, errorResultChan)
}
