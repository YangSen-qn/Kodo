package log

func CalculatePercent(count, totalCount int) float64 {
	if totalCount < 1 {
		return 0
	}
	return float64(count) / float64(totalCount)
}

// --- type
type QueryResultType struct {
	totalCount int
	typeString string
	itemList   []*QueryResultItem
}

func (t *QueryResultType) TotalCount() int {
	return t.totalCount
}

func (t *QueryResultType) TypeString() string {
	return t.typeString
}

func (t *QueryResultType) AllItems() []*QueryResultItem {
	return t.itemList
}

// --- version
type QueryResultVersion struct {
	version       string
	totalCount    int
	successCount  int
	dnsErrorCount int
	typeInfoMap   map[string]*QueryResultType
}

func (version *QueryResultVersion) addTypeInfo(resultType *QueryResultType) {
	if resultType == nil {
		return
	}
	if version.typeInfoMap == nil {
		version.typeInfoMap = make(map[string]*QueryResultType)
	}
	version.typeInfoMap[resultType.typeString] = resultType
}

func (version *QueryResultVersion) calculateCount() {
	version.calculateTotalCount()
	version.calculateSuccessCount()
	version.calculateDnsErrorCount()
}

func (version *QueryResultVersion) calculateTotalCount() {
	count := 0
	typeList := QueryType_Total()
	for _, typeString := range typeList {
		typeInfo := version.typeInfoMap[typeString]
		count += typeInfo.totalCount
	}
	version.totalCount = count
}

func (version *QueryResultVersion) calculateSuccessCount() {
	count := 0
	typeList := QueryType_Success()
	for _, typeString := range typeList {
		typeInfo := version.typeInfoMap[typeString]
		count += typeInfo.totalCount
	}
	version.successCount = count
}

func (version *QueryResultVersion) calculateDnsErrorCount() {
	count := 0
	typeList := QueryType_Dns()
	for _, typeString := range typeList {
		typeInfo := version.typeInfoMap[typeString]
		count += typeInfo.totalCount
	}
	version.dnsErrorCount = count
}

func (version *QueryResultVersion) TotalCount() int {
	return version.totalCount
}

func (version *QueryResultVersion) SuccessCount() int {
	return version.successCount
}

func (version *QueryResultVersion) SuccessPercent() float64 {
	return CalculatePercent(version.successCount, version.totalCount)
}

func (version *QueryResultVersion) DnsErrorCount() int {
	return version.dnsErrorCount
}

func (version *QueryResultVersion) DnsErrorPercent() float64 {
	return CalculatePercent(version.dnsErrorCount, version.totalCount)
}

func (version *QueryResultVersion) Version() string {
	return version.version
}

func (version *QueryResultVersion) TypeInfo(typeString string) *QueryResultType {
	return version.typeInfoMap[typeString]
}

func (version *QueryResultVersion) TypeInfoPercent(typeInfo *QueryResultType) float64 {
	return CalculatePercent(typeInfo.totalCount, version.totalCount)
}
