package uplog

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/YangSen-qn/Kodo/cmd/common"
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/log"
	"github.com/YangSen-qn/Kodo/core/util"

	"github.com/spf13/cobra"
)

type speedPerformer struct {
	config          *common.CommonPerformer
	filePath        string
	fileName        string
	sheet           string
	repoName        string
	startTimeString string
	endTimeString   string
	interval        int64
	queryString     string
	ak              string
	sk              string
}

const (
	_defaultQueryString = `(up_type:"form" OR up_type:"mkblk" OR up_type:"bput" OR up_type:"upload_part") AND status_code:200`
)

func NewSpeedPerformer() *speedPerformer {
	return &speedPerformer{
		config: common.NewPerformer(),
	}
}

func ConfigSpeedCMD(superCMD *cobra.Command) {

	performer := NewSpeedPerformer()

	cmd := &cobra.Command{
		Use:     "speed",
		Short:   "query all type",
		Long:    "",
		Example: "",
		Run:     performer.Execute,
	}

	performer.BindLogCMDToPerformer(cmd)

	superCMD.AddCommand(cmd)
}

func (performer *speedPerformer) BindLogCMDToPerformer(command *cobra.Command) {
	command.Flags().StringVarP(&performer.filePath, "file-path", "", "", "file that result save to")
	command.Flags().StringVarP(&performer.fileName, "file-name", "", "speed.xlsx", "file name that result save to, speed.xlsx")
	command.Flags().StringVarP(&performer.sheet, "sheet", "", "sheet", "sheet that result save to, default sheet")
	command.Flags().StringVarP(&performer.repoName, "repo", "", "", "repo name of query, default use kodo when not set")
	command.Flags().StringVarP(&performer.startTimeString, "start-time", "s", "", "query start time, eg:2020-11-22 00:00:00")
	command.Flags().StringVarP(&performer.endTimeString, "end-time", "e", "", "query end time, eg:2020-11-23 00:00:00")
	command.Flags().StringVarP(&performer.queryString, "query", "", "", "user secret key, default use kodo when not set")
	command.Flags().Int64VarP(&performer.interval, "interval", "", 5*60, "query interval, unit:second, default 5*60s")
	command.Flags().StringVarP(&performer.ak, "ak", "", "", "user access key, default use kodo when not set")
	command.Flags().StringVarP(&performer.sk, "sk", "", "", "user secret key, default use kodo when not set")
}

func (performer *speedPerformer) Execute(cmd *cobra.Command, args []string) {
	// performer.startTimeString = "2021-06-05 00:00:00"
	// performer.endTimeString = "2021-06-08 00:00:00"
	// performer.interval = 5 * 60
	// performer.queryString = "uid:1380337015"

	if performer.startTimeString == "" {
		output.E(errors.New("start time can't empty"))
		return
	}
	if performer.endTimeString == "" {
		output.E(errors.New("end time can't empty"))
		return
	}

	startTime := util.GetTimestampByStringWithDefaultFormat(performer.startTimeString)
	endTime := util.GetTimestampByStringWithDefaultFormat(performer.endTimeString)

	if startTime <= 0 || endTime <= 0 || endTime <= startTime {
		output.E(errors.New("start time or end time is invalid"))
		return
	}

	queryString := performer.queryString
	if len(queryString) == 0 {
		queryString = _defaultQueryString
	} else {
		queryString += " AND " + _defaultQueryString
	}

	output.D().Output("speed param:\n")
	output.D().Output("filePath:" + performer.filePath + "\n")
	output.D().Output("fileName:" + performer.fileName + "\n")
	output.D().Output("sheet   :" + performer.sheet + "\n")
	output.D().Output("start   :" + performer.startTimeString + "\n")
	output.D().Output("end     :" + performer.endTimeString + "\n")
	output.D().Output("query   :" + queryString + "\n")
	output.D().Output("\n")
	output.D().Output("result:\n")

	param := &log.QueryParam{
		RepoName:    performer.repoName,
		StartTime:   startTime,
		EndTime:     endTime,
		QueryString: queryString,
		AK:          performer.ak,
		SK:          performer.sk,
	}

	var speedXlsx *SpeedXlsx
	if performer.filePath != "" {
		path := filepath.Join(performer.filePath, performer.fileName)
		speedXlsx = NewSpeedXlsx(path, performer.sheet)
	}

	log.QuerySpeed(param, performer.interval*1000, func(speedInfo log.QuerySpeedInfo) {
		s := util.GetDateStringWithTimestamp(speedInfo.Start)
		e := util.GetDateStringWithTimestamp(speedInfo.End)
		logInfo := fmt.Sprintf("%v ~ %v speed:%.2fKB/S \n", s, e, speedInfo.Speed)
		output.I().Output(logInfo)

		if speedXlsx != nil {
			if err := speedXlsx.Write(speedInfo); err != nil {
				output.E(errors.New("write err:" + err.Error()))
			}
		}
	})

	if speedXlsx != nil {
		speedXlsx.Compete()
	}
}
