package log

import (
	"fmt"
	"strings"
)

const (
	QS_And = " AND "
	QS_OR  = " OR "

	QS_LogVersion4    = "log_version:4"
	QS_LogTypeQuality = "log_type:quality"
	QS_LogTypeRequest = "log_type:request"

	QS_AndroidDefault = "uid:(NOT \"1380921591\") AND uid:(NOT \"1382167852\") AND uid:(NOT \"1378263724\")"
	QS_iOSDefault     = "uid:(NOT \"1380921591\") AND uid:(NOT \"1380469041\") AND uid:(NOT \"1380333412\")"

	QS_FormatAndroidSDKVersion = "user_agent:QiniuAndroid\\/%s*"
	QS_FormatiOSSDKVersion     = "user_agent:QiniuObject-C\\/%s*"

	QS_Result    = "result"
	QS_ErrorType = "error_type"

	QS_OK            = "ok"
	QS_BadRequest    = "bad_request"
	QS_SSLError      = "ssl_error"
	QS_ResponseError = "response_error"
	QS_FileChanged   = "file_changed"
	QS_ChecksumError = "checksum_error"

	QS_Timeout             = "timeout"
	QS_UnknownHost         = "unknown_host"
	QS_CannotConnectToHost = "cannot_connect_to_host"
	QS_Transmission        = "transmission_error"
	QS_MaliciousResponse   = "malicious_response"
	QS_ParseError          = "parse_error"

	QS_UnknownError           = "unknown_error"
	QS_NetworkError           = "network_error"
	QS_UnexpectedSyscallError = "unexpected_syscall_error"
	QS_LocalIOError           = "local_io_error"
	QS_NetworkSlow            = "network_slow"
	QS_ProtocolError          = "protocol_error"
)

func QueryType_All() []string {
	return []string{
		QS_OK, QS_BadRequest,
		QS_SSLError, QS_ResponseError,
		QS_FileChanged, QS_ChecksumError,
		QS_Timeout, QS_UnknownHost,
		QS_CannotConnectToHost, QS_Transmission,
		QS_MaliciousResponse, QS_ParseError, QS_UnknownError,
		QS_NetworkError, QS_UnexpectedSyscallError, QS_LocalIOError,
		QS_NetworkSlow, QS_ProtocolError,
	}
}

func QueryType_Total() []string {
	return []string{
		QS_OK, QS_BadRequest,
		QS_SSLError, QS_ResponseError,
		QS_FileChanged, QS_ChecksumError,
		QS_Timeout, QS_UnknownHost,
		QS_CannotConnectToHost, QS_Transmission,
		QS_MaliciousResponse, QS_ParseError,
	}
}

func QueryType_Success() []string {
	return []string{
		QS_OK, QS_BadRequest,
		QS_SSLError, QS_ResponseError,
		QS_FileChanged, QS_ChecksumError,
	}
}

func QueryType_Dns() []string {
	return []string{
		QS_UnknownHost,
	}
}

func QueryTypeQualityQueryString(userId string, sdkVersion string, sdkType int, typeList []string) string {
	typeQueryString := queryTypeQueryString(userId, sdkVersion, sdkType, QS_Result, typeList)
	return QS_LogTypeQuality + QS_And + typeQueryString
}

func QueryTypeRequestQueryString(userId string, sdkVersion string, sdkType int, typeList []string) string {
	typeQueryString := queryTypeQueryString(userId, sdkVersion, sdkType, QS_ErrorType, typeList)
	return QS_LogTypeRequest + QS_And + typeQueryString
}

func queryTypeQueryString(userId string, sdkVersion string, sdkType int, typePre string, typeList []string) string {
	var agentAndVersion, defaultContent string
	if sdkType == SDKTypeAndroid {
		defaultContent = QS_AndroidDefault
		agentAndVersion = fmt.Sprintf(QS_FormatAndroidSDKVersion, sdkVersion)
	} else if sdkType == SDKTypeIOS {
		defaultContent = QS_iOSDefault
		agentAndVersion = fmt.Sprintf(QS_FormatiOSSDKVersion, sdkVersion)
	} else if len(sdkVersion) > 0 {
		//"8.1.2,8.2.0" "(sdk_version:8.1.2 OR sdk_version:8.2.0)"
		versionList := strings.Split(sdkVersion, ",")
		defaultContent = strings.Join(versionList, " OR sdk_version:")
		defaultContent = fmt.Sprintf("(sdk_version:%s)", defaultContent)
	}

	typeQueryString := "("
	for i := 0; i < len(typeList); i++ {
		typeQueryString += typePre + ":" + typeList[i]
		if i != len(typeList)-1 {
			typeQueryString += QS_OR
		}
	}
	typeQueryString += ")"

	typeQueryString = QS_LogVersion4 + QS_And + typeQueryString

	if len(userId) > 0 {
		typeQueryString = "uid:" + userId + QS_And + typeQueryString
	}

	if len(agentAndVersion) > 0 {
		typeQueryString += QS_And + agentAndVersion
	}
	if len(defaultContent) > 0 {
		typeQueryString += QS_And + defaultContent
	}

	return typeQueryString
}
