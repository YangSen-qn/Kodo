package uplog

import (
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/log"
	"github.com/YangSen-qn/Kodo/core/util"
	"github.com/spf13/cobra"
	"strings"
)

type networkCMDPerformer struct {
	excelDir        string
	iOSEnable       bool
	androidEnable   bool
	sdkVersion      string
	repoName        string
	startTimeString string
	endTimeString   string
	ak              string
	sk              string
}

func ConfigNetworkSlowCMD(superCMD *cobra.Command) {

	performer := &networkCMDPerformer{}

	cmd := &cobra.Command{
		Use:     "network",
		Short:   "network slow",
		Long:    "",
		Example: "",
		Run:     performer.execute,
	}

	bindNetworkCMDToPerformer(cmd, performer)
	superCMD.AddCommand(cmd)
}

func bindNetworkCMDToPerformer(command *cobra.Command, performer *networkCMDPerformer) {
	command.Flags().StringVarP(&performer.excelDir, "save-file", "o", "", "query result save dir")
	command.Flags().BoolVarP(&performer.iOSEnable, "iOS", "", false, "query iOS result, query iOS And Android if both iOS And Android not set")
	command.Flags().BoolVarP(&performer.androidEnable, "Android", "", false, "query Android result, query iOS And Android if both iOS And Android not set")
	command.Flags().StringVarP(&performer.sdkVersion, "sdk-version", "i", "", "sdk version of query, separate by ',' when have more than one")
	command.Flags().StringVarP(&performer.repoName, "repo", "", "", "repo name of query, default use kodo when not set")
	command.Flags().StringVarP(&performer.startTimeString, "start-time", "s", "", "query start time, eg:2020-11-22 00:00:00")
	command.Flags().StringVarP(&performer.endTimeString, "end-time", "e", "", "query end time, eg:2020-11-23 00:00:00")
	command.Flags().StringVarP(&performer.ak, "ak", "", "", "user access key, default use kodo when not set")
	command.Flags().StringVarP(&performer.sk, "sk", "", "", "user secret key, default use kodo when not set")
}

func (performer *networkCMDPerformer) execute(cmd *cobra.Command, args []string) {
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
		endTime = util.GetCurrentTimestamp()
		startTime = endTime - 5*60*1000

		performer.startTimeString = util.GetDateStringWithTimestamp(startTime)
		performer.endTimeString = util.GetDateStringWithTimestamp(endTime)
	}

	performer.queryByQueryString(startTime, endTime)
}

func (performer *networkCMDPerformer) queryByQueryString(startTime, endTime int64) {
	param := &log.QueryParam{
		SDKType:    0,
		SDKVersion: performer.sdkVersion,
		RepoName:   performer.repoName,
		StartTime:  startTime,
		EndTime:    endTime,
		AK:         performer.ak,
		SK:         performer.sk,
	}

	partResultChan := make(chan *log.QueryResult)
	errorResultChan := make(chan error)

	log.QueryNetworkSlowInfo(param, partResultChan, errorResultChan)

	for partResult := range partResultChan {
		if partResult.AllItems() == nil || len(partResult.AllItems()) == 0 {
			continue
		}
		for _, item := range partResult.AllItems() {
			output.InfoStringFormat("item:", item)
		}
	}
}
