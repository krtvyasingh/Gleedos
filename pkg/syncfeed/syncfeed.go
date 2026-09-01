package syncfeed

type FeedChannel struct {
	FeedURL   string
	LastItem  string
}

func HasNewItem(fc *FeedChannel, latestItem string) bool {
	if fc.LastItem != latestItem {
		fc.LastItem = latestItem
		return true
	}
	return false
}
