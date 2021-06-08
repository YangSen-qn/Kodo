package log

import (
	"fmt"

	"github.com/YangSen-qn/Kodo/core/util"
)

const (
	QueryResultItemRequestId  = "req_id"
	QueryResultItemUserAgent  = "user_agent"
	QueryResultItemUid        = "uid"
	QueryResultItemKeyIP      = "user_ip"
	QueryResultItemHost       = "host"
	QueryResultItemRemoteIP   = "remote_ip"
	QueryResultItemISP        = "isp"
	QueryResultItemCountry    = "country"
	QueryResultItemRegion     = "region"
	QueryResultItemCity       = "city"
	QueryResultItemFileSize   = "file_size"
	QueryResultItemBytesSent  = "bytes_sent"
	QueryResultItemDuration   = "duration"
	QueryResultItemLogVersion = "log_version"
	QueryResultItemStatusCode = "status_code"
	QueryResultItemUpType     = "up_type"
)

type QueryResultItem struct {
	Count            int
	RequestId        string `json:"req_id"`
	UserAgent        string `json:"user_agent"`
	Uid              int    `json:"uid"`
	IP               string `json:"user_ip"`
	Host             string `json:"host"`
	RemoteIP         string `json:"remote_ip"`
	ISP              string `json:"isp"`
	Country          string `json:"country"`
	Region           string `json:"region"`
	City             string `json:"city"`
	FileSize         int64  `json:"file_size"`
	BytesSent        int64  `json:"bytes_sent"`
	Duration         int64  `json:"duration"`
	TotalElapsedTime int64  `json:"total_elapsed_time"`
	LogVersion       int    `json:"log_version"`
	StatusCode       int    `json:"status_code"`
	UpType           string `json:"up_type"`
}

func (item *QueryResultItem) String() string {
	return fmt.Sprintf("{IP:%s, RemoteIP:%s, Host:%s, ISP:%s, Country:%s, Region:%s, City:%s, Count:%d}",
		item.IP, item.RemoteIP, item.Host, item.ISP, item.Country, item.Region, item.City, item.Count)
}

func (item *QueryResultItem) RemoteNetworkType() string {
	if len(item.RemoteIP) > 0 {
		return util.GetIPType(item.RemoteIP)
	} else {
		return item.Host
	}
}

func (item *QueryResultItem) Position() string {
	if len(item.City) > 0 {
		return item.City
	} else if len(item.Region) > 0 {
		return item.Region
	} else {
		return item.Country
	}
}

func (item *QueryResultItem) GetValueByKey(key string) interface{} {
	switch key {
	case QueryResultItemKeyIP:
		return item.IP
	case QueryResultItemHost:
		return item.Host
	case QueryResultItemRemoteIP:
		return item.RemoteIP
	case QueryResultItemISP:
		return item.ISP
	case QueryResultItemCountry:
		return item.Country
	case QueryResultItemRegion:
		return item.Region
	case QueryResultItemCity:
		return item.City
	default:
		return nil
	}
}
