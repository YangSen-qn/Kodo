package log

import (
	"github.com/YangSen-qn/Kodo/core/config"
)

const (
	SDKTypeIOS     = 1
	SDKTypeAndroid = 4
)

type QueryParam struct {
	SDKType     int
	SDKVersion  string
	RepoName    string
	StartTime   int64
	EndTime     int64
	Size        int
	QueryString string

	AK string
	SK string
}

func (info *QueryParam) GetSDKName() string {
	if info.SDKType == SDKTypeAndroid {
		return "Android"
	} else if info.SDKType == SDKTypeIOS {
		return "  iOS  "
	} else {
		return ""
	}
}

func (info *QueryParam) check() {
	if len(info.AK) == 0 {
		info.AK = config.AK
	}
	if len(info.SK) == 0 {
		info.SK = config.SK
	}
	if len(info.RepoName) == 0 {
		info.RepoName = repoName
	}
}

func (info *QueryParam) hasQueryString() bool {
	if len(info.QueryString) == 0 {
		return false
	} else {
		return true
	}
}
