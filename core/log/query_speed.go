package log

type QuerySpeedInfo struct {
	Start int64
	End   int64
	Speed float64
}

type QuerySpeedHandler func(speedInfo QuerySpeedInfo)

func QuerySpeed(param *QueryParam, interval int64, speedHandler QuerySpeedHandler) {
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

		info := QuerySpeedOfDuration(&intervalParam)
		if speedHandler != nil {
			speedHandler(info)
		}

		intervalStart += interval
		intervalEnd += interval
	}
}

func QuerySpeedOfDuration(param *QueryParam) QuerySpeedInfo {

	partResultChan := make(chan *QueryResult)
	errorResultChan := make(chan error)

	speedInfo := QuerySpeedInfo{
		Start: param.StartTime,
		End:   param.EndTime,
		Speed: -1,
	}

	go QueryInfo(param, partResultChan, errorResultChan)

	for result := range partResultChan {
		for _, item := range result.itemList {
			var (
				size     float64 = 0
				duration float64 = 0
			)
			if item.BytesSent > 0 {
				size = float64(item.BytesSent)
			} else if item.FileSize > 0 {
				size = float64(item.FileSize)
			} else {
				continue
			}

			if item.Duration > 0 {
				duration = float64(item.Duration)
			} else {
				continue
			}

			if item.UpType == "jssdk-h5" {
				duration = duration * 1000
			}

			speed := size / duration

			if speedInfo.Speed == -1 {
				speedInfo.Speed = speed
			} else {
				speedInfo.Speed = (speedInfo.Speed + speed) * 0.5
			}
		}
	}

	return speedInfo
}
