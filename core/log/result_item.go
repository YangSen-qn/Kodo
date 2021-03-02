package log

import (
	"fmt"

	"github.com/YangSen-qn/Kodo/core/util"
)

type QueryResultItem struct {
	Count    int
	IP       string
	Host     string
	RemoteIP string
	ISP      string
	Country  string
	Region   string
	City     string
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
	} else if len(item.Region) > 0  {
		return item.Region
	} else {
		return item.Country
	}
}
