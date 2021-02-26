package log

import "fmt"

type QueryResultItem struct {
	Count   int
	IP      string
	ISP     string
	Country string
	Region  string
	City    string
}

func (item *QueryResultItem) String() string {
	return fmt.Sprintf("{IP:%s, ISP:%s, Country:%s, Region:%s, City:%s, Count:%d}",
		item.IP, item.ISP, item.Country, item.Region, item.City, item.Count)
}
