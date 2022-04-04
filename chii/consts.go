package chii

const bgmAPI = "https://api.bgm.tv/v0/"

type CollectionType string

const (
	CollectionTypeWatching    CollectionType = "watching"     // 在看的动画与三次元条目
	CollectionTypeAllWatching CollectionType = "all_watching" // 在看的动画三次元与书籍条目
)

type CollectionsStatus string

const (
	CollectionsStatusWish    CollectionsStatus = "wish"
	CollectionsStatusCollect CollectionsStatus = "collect"
	CollectionsStatusDo      CollectionsStatus = "do"
	CollectionsStatusOnHold  CollectionsStatus = "on_hold"
	CollectionsStatusDropped CollectionsStatus = "dropped"
)

type ResponseGroup string

const (
	ResponseSmall  ResponseGroup = "small"
	ResponseMediam ResponseGroup = "mediam"
	ResponseLarge  ResponseGroup = "large"
)

var StatusCode = map[int]string{
	200:   "OK",
	202:   "Accepted",
	304:   "Not Modified",
	30401: "Not Modified: Collection already exists",

	400:   "Bad Request",
	40001: "Error: Nothing found with that ID",

	401:   "Unauthorized",
	40101: "Error: Auth failed over 5 times",
	40102: "Error: Username is not an Email address",

	404: "Not Found",
	405: "Method Not Allowed",
}
