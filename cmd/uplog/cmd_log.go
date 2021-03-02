package uplog

import (
	"github.com/YangSen-qn/Kodo/cmd/common"
	"path/filepath"
	"strings"

	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/log"
	"github.com/YangSen-qn/Kodo/core/util"
	"github.com/spf13/cobra"
)

type upLogPerformer struct {
	config          *common.CommonPerformer
	excelDir        string
	iOSEnable       bool
	androidEnable   bool
	sdkVersion      string
	repoName        string
	beforeMinute    int64
	startTimeString string
	endTimeString   string
	queryString     string
	ak              string
	sk              string
}

func NewIPPerformer() *upLogPerformer {
	return &upLogPerformer{
		config: common.NewPerformer(),
	}
}

func ConfigCMD(superCMD *cobra.Command) {

	performer := NewIPPerformer()

	cmd := &cobra.Command{
		Use:     "uplog",
		Short:   "query uplog",
		Long:    "",
		Example: "",
		Run:     performer.Execute,
	}

	performer.BindLogCMDToPerformer(cmd)

	superCMD.AddCommand(cmd)
}

func (performer *upLogPerformer) BindLogCMDToPerformer(command *cobra.Command) {
	command.Flags().StringVarP(&performer.excelDir, "save-file", "o", "", "query result save dir")
	command.Flags().BoolVarP(&performer.iOSEnable, "iOS", "", false, "query iOS result, query iOS And Android if both iOS And Android not set")
	command.Flags().BoolVarP(&performer.androidEnable, "Android", "", false, "query Android result, query iOS And Android if both iOS And Android not set")
	command.Flags().StringVarP(&performer.sdkVersion, "sdk-version", "i", "", "sdk version of query, separate by ',' when have more than one")
	command.Flags().StringVarP(&performer.repoName, "repo", "", "", "repo name of query, default use kodo when not set")
	command.Flags().Int64VarP(&performer.beforeMinute, "before", "b", 5, "query minutes before current time, default 5, when start-time or end-time was not set use this")
	command.Flags().StringVarP(&performer.startTimeString, "start-time", "s", "", "query start time, eg:2020-11-22 00:00:00")
	command.Flags().StringVarP(&performer.endTimeString, "end-time", "e", "", "query end time, eg:2020-11-23 00:00:00")
	command.Flags().StringVarP(&performer.queryString, "query-info", "", "", "query info string, when set sdk-version iOS Android and save-file will invalid")
	command.Flags().StringVarP(&performer.ak, "ak", "", "", "user access key, default use kodo when not set")
	command.Flags().StringVarP(&performer.sk, "sk", "", "", "user secret key, default use kodo when not set")
}

func (performer *upLogPerformer) Execute(cmd *cobra.Command, args []string) {
	performer.sdkVersion = strings.ReplaceAll(performer.sdkVersion, " ", "")
	if len(performer.sdkVersion) == 0 {
		if performer.iOSEnable && performer.androidEnable {
			performer.sdkVersion = iOSAndAndroidAllVersion
		} else if performer.iOSEnable {
			performer.sdkVersion = iOSAllVersion
		} else if performer.androidEnable {
			performer.sdkVersion = androidAllVersion
		}
	}

	var startTime, endTime int64
	if len(performer.startTimeString) > 0 && len(performer.endTimeString) > 0 {
		startTime = util.GetTimestampByStringWithDefaultFormat(performer.startTimeString)
		endTime = util.GetTimestampByStringWithDefaultFormat(performer.endTimeString)
	}

	if startTime <= 0 || endTime <= 0 {
		if performer.beforeMinute <= 0 {
			performer.beforeMinute = 5
		}
		endTime = util.GetCurrentTimestamp()
		startTime = endTime - int64(performer.beforeMinute)*60*1000

		performer.startTimeString = util.GetDateStringWithTimestamp(startTime)
		performer.endTimeString = util.GetDateStringWithTimestamp(endTime)
	}

	if len(performer.queryString) != 0 {
		performer.queryByQueryString(startTime, endTime)
	} else {
		performer.queryVersions(startTime, endTime)
	}
}

func (performer *upLogPerformer) queryByQueryString(startTime, endTime int64) {
	param := &log.QueryParam{
		SDKType:     0,
		SDKVersion:  performer.sdkVersion,
		RepoName:    performer.repoName,
		StartTime:   startTime,
		EndTime:     endTime,
		QueryString: performer.queryString,
		AK:          performer.ak,
		SK:          performer.sk,
	}

	result, err := log.QueryByQueryString(param)

	if err != nil {
		output.E(err)
		return
	}

	output.I().OutputFormat("count:%d", result.TotalCount())
}

// query version
func (performer *upLogPerformer) queryVersions(startTime, endTime int64) {
	var versionArray []string
	if strings.Contains(performer.sdkVersion, ",") {
		versionArray = strings.Split(performer.sdkVersion, ",")
	} else {
		versionArray = []string{performer.sdkVersion}
	}

	queryTypeArray := []int{}
	if performer.iOSEnable {
		queryTypeArray = append(queryTypeArray, log.SDKTypeIOS)
	}
	if performer.androidEnable {
		queryTypeArray = append(queryTypeArray, log.SDKTypeAndroid)
	}
	if len(queryTypeArray) == 0 {
		queryTypeArray = append(queryTypeArray, log.SDKTypeIOS, log.SDKTypeAndroid)
	}

	for _, queryType := range queryTypeArray {
		paramList := []*log.QueryParam{}
		for _, version := range versionArray {
			param := &log.QueryParam{
				SDKType:    queryType,
				SDKVersion: version,
				RepoName:   performer.repoName,
				StartTime:  startTime,
				EndTime:    endTime,
				AK:         performer.ak,
				SK:         performer.sk,
			}
			paramList = append(paramList, param)
		}
		performer.querySomeVersion(paramList, log.QueryType_All())
	}
}

func (performer *upLogPerformer) querySomeVersion(paramList []*log.QueryParam, types []string) {

	if paramList == nil || len(paramList) == 0 {
		return
	}

	var allVersionLogCount int = 0
	resultVersionList := [] *log.QueryResultVersion{}
	for _, param := range paramList {
		resultVersion := log.QueryVersion(param, types)
		allVersionLogCount += resultVersion.TotalCount()
		resultVersionList = append(resultVersionList, resultVersion)

		outputVersion(param.GetSDKName(), resultVersion)
	}

	outputVersionList("  整体 信息  ", allVersionLogCount, resultVersionList)

	if performer.excelDir != "" {
		fileName := performer.startTimeString + "~" + performer.endTimeString + ".xlsx"
		fileName = strings.ReplaceAll(fileName, "-", "")
		fileName = strings.ReplaceAll(fileName, " ", "")
		fileName = strings.ReplaceAll(fileName, ":", "")
		saveVersionToLocalAsExcel(filepath.Join(performer.excelDir, fileName), paramList[0].GetSDKName(), allVersionLogCount, resultVersionList, types)
	}
}
