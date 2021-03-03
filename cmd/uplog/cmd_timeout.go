package uplog

import (
	"github.com/YangSen-qn/Kodo/cmd/common"
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/log"
	"github.com/YangSen-qn/Kodo/core/util"
	"github.com/spf13/cobra"
	"path/filepath"
	"sort"
	"strings"
)

type timeoutCMDPerformer struct {
	config          *common.CommonPerformer
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

type timeoutResultGroup struct {
	count int
	group map[string]*log.QueryResultItem
}

func NewTimeoutPerformer() *timeoutCMDPerformer {
	return &timeoutCMDPerformer{
		config: common.NewPerformer(),
	}
}

func ConfigTimeoutCMD(superCMD *cobra.Command) {

	performer := NewTimeoutPerformer()

	cmd := &cobra.Command{
		Use:     "timeout",
		Short:   "get timeout error info",
		Long:    "",
		Example: "",
		Run:     performer.Execute,
	}

	performer.config.BindToCMD(cmd)
	performer.BindToCMD(cmd)
	superCMD.AddCommand(cmd)
}

func (performer *timeoutCMDPerformer) BindToCMD(command *cobra.Command) {
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

func (performer *timeoutCMDPerformer) Execute(cmd *cobra.Command, args []string) {
	performer.sdkVersion = strings.ReplaceAll(performer.sdkVersion, " ", "")
	if len(performer.sdkVersion) == 0 {
		performer.sdkVersion = allTimeoutVersion
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

func (performer *timeoutCMDPerformer) queryByQueryString(startTime, endTime int64) {
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

	go log.QueryInfoOfCannotConnectToServer(param, partResultChan, errorResultChan)

	excelRowCount := 0 // excel 行数
	index := 0
	result := make(map[string]*timeoutResultGroup)
	for partResult := range partResultChan {
		if partResult.AllItems() == nil || len(partResult.AllItems()) == 0 {
			continue
		}
		for _, item := range partResult.AllItems() {

			itemKey := item.Position()
			if len(itemKey) > 0 {
				group := result[itemKey]
				if group == nil {
					group = &timeoutResultGroup{
						count: 0,
						group: make(map[string]*log.QueryResultItem),
					}
					result[itemKey] = group
				}

				itemRemoteNetworkType := item.RemoteNetworkType()
				oldItem := group.group[itemRemoteNetworkType]
				if oldItem == nil {
					excelRowCount++
					group.count++
					group.group[itemRemoteNetworkType] = item
				} else {
					group.count++
					oldItem.Count += 1
				}
			}

			index++
			output.W().OutputFormat("item:%s index:%d/total:%d\n", item, index, partResult.TotalCount())
		}

	}

	if len(performer.excelDir) > 0 {
		items := make([]*log.QueryResultItem, 0, excelRowCount)

		// 根据 Position 对应所有 item 数量进行排序
		groupList := make([]*timeoutResultGroup, 0, len(result))
		for _, group := range result {
			groupList = append(groupList, group)
		}
		sort.Slice(groupList, func(i, j int) bool {
			return groupList[i].count > groupList[j].count
		})

		// 从排完序的 group 中读取 items 依次排序并加入 items 中
		for _, group := range groupList{

			// 获取 groupItems 并排序
			groupItems := make([]*log.QueryResultItem, 0, len(group.group))
			for _, item := range group.group{
				groupItems = append(groupItems, item)
			}
			sort.Slice(groupItems, func(i, j int) bool {
				return groupItems[i].Count > groupItems[j].Count
			})

			// 添加到 items 中
			items = append(items, groupItems...)
		}

		err := saveResultItemsToLocalAsExcel(filepath.Join(performer.excelDir, "timeout.xlsx"), items)
		if err != nil {
			output.I().OutputFormat("save error:%s", err)
		}
	}
}
