package log

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/qiniu/pandora-go-sdk/base"
	. "github.com/qiniu/pandora-go-sdk/logdb"
)

func queryCountSeparateByDuration(param *QueryParam) (result *QueryResult, err error) {
	partQueryInfo := *param

	result = &QueryResult{
		totalCount: 0,
		itemList:   nil,
	}

	var partEndTime int64 = 0
	for partStartTime := param.StartTime; partEndTime < param.EndTime; {
		partQueryInfo := partQueryInfo

		partEndTime = partStartTime + 4*3600*1000
		if partEndTime > param.EndTime {
			partEndTime = param.EndTime
		}

		partQueryInfo.StartTime = partStartTime
		partQueryInfo.EndTime = partEndTime
		partResult, e := queryCountByParam(partQueryInfo)
		err = e
		if partResult != nil {
			result.addResult(partResult)
		}

		partStartTime = partEndTime
	}

	return
}

func queryCountByParam(param QueryParam) (result *QueryResult, err error) {

	cfg := NewConfig().
		WithEndpoint(logDBEndPoint).
		WithAccessKeySecretKey(param.AK, param.SK).
		WithLogger(NewDefaultLogger()).
		WithLoggerLevel(LogDebug).
		WithDialTimeout(180 * time.Second).
		WithResponseTimeout(180 * time.Second)

	client, err := New(cfg);
	if err != nil {
		return
	}

	logInput := &PartialQueryInput{
		RepoName:    param.RepoName,
		StartTime:   param.StartTime,
		EndTime:     param.EndTime,
		QueryString: param.QueryString,
		SearchType:  PartialQuerySearchTypeA,
		Size:        1,
		Sort:        "up_time",
	}

	logOutput, err := client.PartialQuery(logInput)
	if err != nil {
		return
	}

	result = &QueryResult{
		totalCount: logOutput.Total,
		itemList:   nil,
	}

	return
}

func queryInfoSeparateByPage(param *QueryParam, partResultChan chan<- *QueryResult, errorResultChan chan<- error) {
	partIndex := 0
	for {
		partIndex++
		result, err := queryPartInfoByParam(partIndex, "10m", param)
		if result != nil {
			partResultChan <- result
		}
		if err != nil {
			errorResultChan <- err
		}
	}
	return
}

func queryPartInfoByParam(partIndex int, scroll string, param *QueryParam) (result *QueryResult, err error) {

	cfg := NewConfig().
		WithEndpoint(logDBEndPoint).
		WithAccessKeySecretKey(param.AK, param.SK).
		WithLogger(NewDefaultLogger()).
		WithLoggerLevel(LogDebug).
		WithDialTimeout(180 * time.Second).
		WithResponseTimeout(180 * time.Second)

	client, err := New(cfg);
	if err != nil {
		return
	}

	startTime := strconv.Itoa(int(param.StartTime))
	endTime := strconv.Itoa(int(param.EndTime))
	queryString := fmt.Sprintf("up_time:>%s AND up_time:<%s AND %s", startTime, endTime, param.QueryString)
	logInput := &QueryLogInput{
		RepoName:  param.RepoName,
		Query:     queryString,
		Sort:      "up_time:asc",
		From:      partIndex,
		Size:      10,
		Scroll:    scroll,
		Highlight: nil,
	}
	logOutput, err := client.QueryLog(logInput)

	fmt.Println("query string:", queryString)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(logOutput)

	var itemList []*QueryResultItem = nil
	if len(logOutput.Data) > 0 {
		itemList = make([]*QueryResultItem, logOutput.Total)

	}

	result = &QueryResult{
		totalCount: logOutput.Total,
		itemList:   itemList,
	}

	return
}
