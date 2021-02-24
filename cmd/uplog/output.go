package uplog

import (
	"github.com/YangSen-qn/Kodo/cmd/output"
	"github.com/YangSen-qn/Kodo/core/log"
	"strconv"
)

func outputTitleInfo(title string) {
	style := output.NewOutputMessageStyle().Color(output.PrintForegroundColorGreen)
	output.InfoStringFormatWithStyle(style, title, ":\n")
}

func outputVersionTitle(title string) {
	style := output.NewOutputMessageStyle().Color(output.PrintForegroundColorGreen)
	output.InfoStringFormatWithStyle(style, "============================= "+title+" =============================\n")
}

func outputContentInfo(content string) {
	style := output.NewOutputMessageStyle().Color(output.PrintForegroundColorYellow)
	output.InfoStringFormatWithStyle(style, content, "\n")
}

func outputLogResult(result string, count int, percent float64) {
	resultLabelStyle := output.NewOutputMessageStyle().Width(40).Color(output.PrintForegroundColorWhite)
	resultContentStyle := output.NewOutputMessageStyle().Width(15).Color(output.PrintForegroundColorYellow)
	percentLabelStyle := output.NewOutputMessageStyle().Width(9).Color(output.PrintForegroundColorWhite)
	percentContentStyle := output.NewOutputMessageStyle().Width(9).Color(output.PrintForegroundColorYellow)

	output.InfoStringFormatWithStyle(resultLabelStyle, result+":")
	output.InfoStringFormatWithStyle(resultContentStyle, strconv.Itoa(count))

	if percent >= 0 {
		output.InfoStringFormatWithStyle(percentLabelStyle, "percent:")
		output.InfoStringFormatWithStyle(percentContentStyle, "%.4f%%", percent*100)
	}
	output.InfoStringFormat("\n")
}

// 输出多个 Version 整体信息
func outputVersion(sdkName string, version *log.QueryResultVersion) {
	if version == nil {
		return
	}

	outputVersionTitle(sdkName + " " + version.Version())
	outputLogResult(version.Version()+" totalCount", version.TotalCount(), -1)
	outputLogResult(version.Version()+" successCount", version.SuccessCount(), version.SuccessPercent())
	outputLogResult(version.Version()+" dnsError", version.DnsErrorCount(), version.DnsErrorPercent())

	allTypes := log.QueryType_All()
	for _, t := range allTypes {
		typeResult := version.TypeInfo(t)
		typePercent := version.TypeInfoPercent(typeResult)
		outputLogResult(version.Version()+" "+t, typeResult.TotalCount(), typePercent)
	}
}

func outputVersionList(sdkName string, allVersionLogCount int, versionLogInfoList [] *log.QueryResultVersion) {
	if versionLogInfoList == nil {
		return
	}
	outputVersionTitle(sdkName)
	outputLogResult("total", allVersionLogCount, -1)
	for _, version := range versionLogInfoList {
		outputLogResult(version.Version(), version.TotalCount(), float64(version.TotalCount())/float64(allVersionLogCount))
	}
}
