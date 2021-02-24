package log

import (
	"time"

	. "github.com/qiniu/pandora-go-sdk/base"
	. "github.com/qiniu/pandora-go-sdk/logdb"
)

func queryByPart(param *QueryParam) (result *QueryResult, err error) {

	result = &QueryResult{
		totalCount: 0,
		itemList:   nil,
	}

	var partEndTime int64 = 0
	for partStartTime := param.StartTime; partEndTime < param.EndTime; {
		partQueryInfo := param

		partEndTime = partStartTime + 4*3600*1000
		if partEndTime > param.EndTime {
			partEndTime = param.EndTime
		}

		partQueryInfo.StartTime = partStartTime
		partQueryInfo.EndTime = partEndTime
		partResult, e := queryByParam(param)
		err = e
		result.addResult(partResult)

		partStartTime = partEndTime
	}

	return
}

func queryByParam(param *QueryParam) (result *QueryResult, err error) {

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
