package util

import (
	"time"
)

func utilGetTimestampByStringWithDefaultFormat(timeString string) int64 {

	return utilGetTimestampByString(timeString, "2006-01-02 03:04:05")
}

func utilGetTimestampByString(timeString string, format string) int64 {
	cstZone := time.FixedZone("GMT", 8*3600)
	tm, err := time.ParseInLocation(format, timeString, cstZone)
	if err != nil {
		return -1
	} else {
		return tm.Local().Unix() * 1000
	}
}

func utilGetCurrentTimestamp() int64 {
	timeUnix := time.Now().Unix()
	return timeUnix * 1000
}