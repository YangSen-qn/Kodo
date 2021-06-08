package uplog

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/YangSen-qn/Kodo/cmd/common"
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/log"
	"github.com/YangSen-qn/Kodo/core/util"

	"github.com/spf13/cobra"
)

type allTypePerformer struct {
	config          *common.CommonPerformer
	repoName        string
	startTimeString string
	endTimeString   string
	typeKeyList     string
	queryString     string
	ak              string
	sk              string
}

func NewAllTypePerformer() *upLogPerformer {
	return &upLogPerformer{
		config: common.NewPerformer(),
	}
}

func ConfigAllTypeCMD(superCMD *cobra.Command) {

	performer := NewIPPerformer()

	cmd := &cobra.Command{
		Use:     "allType",
		Short:   "query all type",
		Long:    "",
		Example: "",
		Run:     performer.Execute,
	}

	performer.BindLogCMDToPerformer(cmd)

	superCMD.AddCommand(cmd)
}

func (performer *allTypePerformer) BindLogCMDToPerformer(command *cobra.Command) {
	command.Flags().StringVarP(&performer.repoName, "repo", "", "", "repo name of query, default use kodo when not set")
	command.Flags().StringVarP(&performer.startTimeString, "start-time", "s", "", "query start time, eg:2020-11-22 00:00:00")
	command.Flags().StringVarP(&performer.endTimeString, "end-time", "e", "", "query end time, eg:2020-11-23 00:00:00")
	command.Flags().StringVarP(&performer.queryString, "query", "", "", "user secret key, default use kodo when not set")
	command.Flags().StringVarP(&performer.typeKeyList, "typeList", "", "", "type list, eg:[{"key":"version", "region":{"location":0,"length":2}}]")
	command.Flags().StringVarP(&performer.ak, "ak", "", "", "user access key, default use kodo when not set")
	command.Flags().StringVarP(&performer.sk, "sk", "", "", "user secret key, default use kodo when not set")
}

func (performer *allTypePerformer) Execute(cmd *cobra.Command, args []string) {
	if performer.startTimeString == "" {
		output.E(errors.New("start time can't empty"))
		return
	}
	if performer.endTimeString == "" {
		output.E(errors.New("end time can't empty"))
		return
	}
	if performer.typeKeyList == "" {
		output.E(errors.New("type list can't empty"))
		return
	}

	var typeParamList []*log.TypeParam
	if err := json.Unmarshal([]byte(performer.typeKeyList), &typeParamList); err != nil {
		output.E(errors.New("type list err:" + err.Error()))
		return
	}

	startTime := util.GetTimestampByStringWithDefaultFormat(performer.startTimeString)
	endTime := util.GetTimestampByStringWithDefaultFormat(performer.endTimeString)

	if startTime <= 0 || endTime <= 0 || endTime <= startTime {
		output.E(errors.New("start time or end time is invalid"))
		return
	}

	param := &log.QueryParam{
		StartTime:   startTime,
		EndTime:     endTime,
		QueryString: performer.queryString,
	}

	allType := log.QueryAllType(typeParamList, param)
	for key, value := range allType {
		output.I().Output("key:" + key)
		for _, item := range value {
			output.I().Output("type:" + item.TypeString + " count:" + strconv.Itoa(item.Count))
		}
	}
}
