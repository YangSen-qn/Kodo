package log

import (
	"sync"
)

const (
	_allQueryString         = `(up_type:"form" OR up_type:"mkfile" OR up_type:"jssdk-h5")`
	_successQueryString     = _allQueryString + ` AND status_code:200`
	_serverErrorQueryString = _allQueryString + ` AND status_code:>299`
)

type QuerySuccessRateInfo struct {
	Start            int64
	End              int64
	AllCount         int64
	SuccessCount     int64
	ServerErrorCount int64
}

type QuerySuccessRateHandler func(successRateInfo QuerySuccessRateInfo)

func QuerySuccessRate(param *QueryParam, interval int64, successRateHandler QuerySuccessRateHandler) {
	if param == nil {
		param = &QueryParam{}
	}
	param.check()

	intervalStart := param.StartTime
	intervalEnd := intervalStart + interval
	for intervalStart < param.EndTime {
		if intervalEnd > param.EndTime {
			intervalEnd = param.EndTime
		}
		intervalParam := *param
		intervalParam.StartTime = intervalStart
		intervalParam.EndTime = intervalEnd

		info := QuerySuccessRateOfDuration(&intervalParam)
		if successRateHandler != nil {
			successRateHandler(info)
		}

		intervalStart += interval
		intervalEnd += interval
	}
}

func QuerySuccessRateOfDuration(param *QueryParam) QuerySuccessRateInfo {

	info := QuerySuccessRateInfo{
		Start:            param.StartTime,
		End:              param.EndTime,
		AllCount:         0,
		SuccessCount:     0,
		ServerErrorCount: 0,
	}

	w := &sync.WaitGroup{}
	w.Add(3)
	go func() {
		p := *param
		p.QueryString += " AND " + _allQueryString
		result, _ := queryCountByParam(p)
		info.AllCount = int64(result.TotalCount())
		w.Done()
	}()

	go func() {
		p := *param
		p.QueryString += " AND " + _successQueryString
		result, _ := queryCountByParam(p)
		info.SuccessCount = int64(result.TotalCount())
		w.Done()
	}()

	go func() {
		p := *param
		p.QueryString += " AND " + _serverErrorQueryString
		result, _ := queryCountByParam(p)
		info.ServerErrorCount = int64(result.TotalCount())
		w.Done()
	}()

	w.Wait()

	return info
}
