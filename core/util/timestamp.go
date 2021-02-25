package util

import (
	"time"
)

func GetDateStringWithTimestamp(timestamp int64) string {
	t := time.Unix(timestamp/1000, 0)
	return t.Format("2006-01-02 03:04:05")
}

func GetTimestampByStringWithDefaultFormat(timeString string) int64 {
	return GetTimestampByString(timeString, "2006-01-02 15:04:05")
}

func GetTimestampByString(timeString string, format string) int64 {
	cstZone := time.FixedZone("GMT", 8*3600)
	tm, err := time.ParseInLocation(format, timeString, cstZone)
	if err != nil {
		return -1
	} else {
		return tm.Local().Unix() * 1000
	}
}

func GetCurrentTimestamp() int64 {
	timeUnix := time.Now().Unix()
	return timeUnix * 1000
}