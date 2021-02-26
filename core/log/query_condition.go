package log


import (
	"fmt"
)

const (

	QS_And = " AND "
	QS_OR  = " OR "

	QS_LogVersion4    = "log_version:4"
	QS_LogTypeQuality = "log_type:quality"

	QS_AndroidDefault = "uid:(NOT \"1380921591\") AND uid:(NOT \"1382167852\") AND uid:(NOT \"1378263724\")"
	QS_iOSDefault     = "uid:(NOT \"1380921591\") AND uid:(NOT \"1380469041\") AND uid:(NOT \"1380333412\")"

	QS_FormatAndroidSDKVersion = "user_agent:QiniuAndroid\\/%s*"
	QS_FormatiOSSDKVersion     = "user_agent:QiniuObject-C\\/%s*"

	QS_ResultOK            = "result:ok"
	QS_ResultBadRequest    = "result:bad_request"
	QS_ResultSSLError      = "result:ssl_error"
	QS_ResultResponseError = "result:response_error"
	QS_ResultFileChanged   = "result:file_changed"
	QS_ResultChecksumError = "result:checksum_error"

	QS_ResultTimeout             = "result:timeout"
	QS_ResultUnknownHost         = "result:unknown_host"
	QS_ResultCannotConnectToHost = "result:cannot_connect_to_host"
	QS_ResultTransmission        = "result:transmission_error"
	QS_ResultMaliciousResponse   = "result:malicious_response"
	QS_ResultParseError          = "result:parse_error"

	QS_ResultUnknownError           = "result:unknown_error"
	QS_ResultNetworkError           = "result:network_error"
	QS_ResultUnexpectedSyscallError = "result:unexpected_syscall_error"
	QS_ResultLocalIOError           = "result:local_io_error"
	QS_ResultNetworkSlow            = "result:network_slow"
	QS_ResultProtocolError          = "result:protocol_error"
)

func QueryType_All() []string {
	return []string{
		QS_ResultOK, QS_ResultBadRequest,
		QS_ResultSSLError, QS_ResultResponseError,
		QS_ResultFileChanged, QS_ResultChecksumError,
		QS_ResultTimeout, QS_ResultUnknownHost,
		QS_ResultCannotConnectToHost, QS_ResultTransmission,
		QS_ResultMaliciousResponse, QS_ResultParseError, QS_ResultUnknownError,
		QS_ResultNetworkError, QS_ResultUnexpectedSyscallError, QS_ResultLocalIOError,
		QS_ResultNetworkSlow, QS_ResultProtocolError,
	}
}

func QueryType_Total() []string {
	return []string{
		QS_ResultOK, QS_ResultBadRequest,
		QS_ResultSSLError, QS_ResultResponseError,
		QS_ResultFileChanged, QS_ResultChecksumError,
		QS_ResultTimeout, QS_ResultUnknownHost,
		QS_ResultCannotConnectToHost, QS_ResultTransmission,
		QS_ResultMaliciousResponse, QS_ResultParseError,
	}
}

func QueryType_Success() []string {
	return []string{
		QS_ResultOK, QS_ResultBadRequest,
		QS_ResultSSLError, QS_ResultResponseError,
		QS_ResultFileChanged, QS_ResultChecksumError,
	}
}

func QueryType_Dns() []string {
	return []string{
		QS_ResultUnknownHost,
	}
}

func QueryTypeQueryString(sdkVersion string, sdkType int, typeList []string) string {
	var agentAndVersion, defaultContent string
	if sdkType == SDKTypeAndroid {
		defaultContent = QS_AndroidDefault
		agentAndVersion = fmt.Sprintf(QS_FormatAndroidSDKVersion, sdkVersion)
	} else if sdkType == SDKTypeIOS {
		defaultContent = QS_iOSDefault
		agentAndVersion = fmt.Sprintf(QS_FormatiOSSDKVersion, sdkVersion)
	} else {
		defaultContent = "sdk_version:" + sdkVersion
	}

	typeQueryString := "("
	for i := 0; i < len(typeList); i++ {
		typeQueryString += typeList[i]
		if i != len(typeList)-1 {
			typeQueryString += QS_OR
		}
	}
	typeQueryString += ")"

	typeQueryString = QS_LogVersion4 + QS_And + QS_LogTypeQuality + QS_And + typeQueryString

	if len(agentAndVersion) > 0 {
		typeQueryString += QS_And + agentAndVersion
	}
	if len(defaultContent) > 0 {
		typeQueryString += QS_And + defaultContent
	}

	return typeQueryString
}