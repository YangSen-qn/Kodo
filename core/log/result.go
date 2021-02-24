package log

import "sync"

type QueryResult struct {
	lock       sync.Mutex
	totalCount int
	itemList   []*QueryResultItem
}

func (result *QueryResult) TotalCount() int {
	return result.totalCount
}

func (result *QueryResult) AllItems() []*QueryResultItem {
	return result.itemList
}

func (result *QueryResult) addResult(newResult *QueryResult) {
	result.lock.Lock()
	result.totalCount += newResult.totalCount
	if result.itemList != nil && newResult.itemList != nil && len(newResult.itemList) > 0 {
		result.itemList = append(result.itemList, newResult.itemList...)
	}
	result.lock.Unlock()
}
