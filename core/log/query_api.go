package log

import (
	"fmt"
	"github.com/qiniu/pandora-go-sdk/base/config"
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
	defer func() {
		close(partResultChan)
		close(errorResultChan)
	}()

	startTime := strconv.Itoa(int(param.StartTime))
	endTime := strconv.Itoa(int(param.EndTime))
	param.QueryString = fmt.Sprintf("up_time:>%s AND up_time:<%s AND %s", startTime, endTime, param.QueryString)

	config := NewConfig().
		WithEndpoint(logDBEndPoint).
		WithAccessKeySecretKey(param.AK, param.SK).
		WithLogger(NewDefaultLogger()).
		WithLoggerLevel(LogDebug).
		WithDialTimeout(180 * time.Second).
		WithResponseTimeout(180 * time.Second)

	count := 0
	scrollId, partResult, err := queryPartInfoInitByParam(param, config)
	if partResult != nil {
		if partResult.itemList != nil && len(partResult.itemList) > 0 {
			count += len(partResult.AllItems())
			partResultChan <- partResult
		}
	}

	for {
		scrollId, partResult, err = queryLeftPartInfoByParam(scrollId, param, config)
		if partResult != nil {
			if partResult.itemList == nil || len(partResult.itemList) == 0 {
				break
			}
			count += len(partResult.AllItems())
			partResultChan <- partResult
		}
		if err != nil {
			errorResultChan <- err
			break
		}
		if count >= partResult.totalCount {
			break
		}
	}
	return
}

func queryPartInfoInitByParam(param *QueryParam, config *config.Config) (scrollId string, result *QueryResult, err error) {
	client, err := New(config);
	if err != nil {
		return
	}

	logInput := &QueryLogInput{
		RepoName:  param.RepoName,
		Query:     param.QueryString,
		Sort:      "up_time:asc",
		From:      0,
		Size:      20,
		Scroll:    "10m",
		Highlight: nil,
	}
	logOutput, err := client.QueryLog(logInput)

	scrollId = logOutput.ScrollId
	result = queryInfoOutputToResult(logOutput)

	return
}

func queryLeftPartInfoByParam(scrollId string, param *QueryParam, config *config.Config) (newScrollId string, result *QueryResult, err error) {
	client, err := New(config);
	if err != nil {
		return
	}

	logInput := &QueryScrollInput{
		RepoName: param.RepoName,
		Scroll:   "10m",
		ScrollId: scrollId,
	}
	logOutput, err := client.QueryScroll(logInput)

	newScrollId = logOutput.ScrollId
	result = queryInfoOutputToResult(logOutput)

	return
}

func queryInfoOutputToResult(logOutput *QueryLogOutput) *QueryResult {
	var itemList []*QueryResultItem = nil
	if len(logOutput.Data) > 0 {
		itemList = make([]*QueryResultItem, 0, len(logOutput.Data))
		for _, itemData := range logOutput.Data {
			item := &QueryResultItem{Count:1}
			if itemData["user_ip"] != nil {
				item.IP = itemData["user_ip"].(string)
			}
			if itemData["isp"] != nil {
				item.ISP = itemData["isp"].(string)
			}
			if itemData["country"] != nil {
				item.Country = itemData["country"].(string)
			}
			if itemData["region"] != nil {
				item.Region = itemData["region"].(string)
			}
			if itemData["city"] != nil {
				item.City = itemData["city"].(string)
			}
			if itemData["host"] != nil {
				item.Host = itemData["host"].(string)
			}
			if itemData["remote_ip"] != nil {
				item.RemoteIP = itemData["remote_ip"].(string)
			}
			itemList = append(itemList, item)
		}
	}

	return &QueryResult{
		totalCount: logOutput.Total,
		itemList:   itemList,
	}
}
