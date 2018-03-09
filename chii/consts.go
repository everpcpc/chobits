package chii

const bgmAPI = "https://api.bgm.tv/"

// SubjectType ...
type SubjectType int

// ...
const (
	TypeNil SubjectType = iota
	TypeBook
	TypeAnime
	TypeMusic
	TypeGame
	TypeReserved
	TypeReal
)

var subjectTypeString = map[SubjectType]string{
	TypeNil:      "nil",
	TypeBook:     "book",
	TypeAnime:    "anime",
	TypeMusic:    "music",
	TypeGame:     "game",
	TypeReserved: "reserved",
	TypeReal:     "real",
}

// CollectionType ...
type CollectionType string

// ...
const (
	TypeWatching    CollectionType = "watching"     // 在看的动画与三次元条目
	TypeAllWatching CollectionType = "all_watching" // 在看的动画三次元与书籍条目
)

// CollectionsStatus ...
type CollectionsStatus string

// ...
const (
	StatusWish    CollectionsStatus = "wish"
	StatusCollect CollectionsStatus = "collect"
	StatusDo      CollectionsStatus = "do"
	StatusOnHold  CollectionsStatus = "on_hold"
	StatusDropped CollectionsStatus = "dropped"
)

// EpStatus ...
type EpStatus string

// ...
const (
	StatusEpWatched EpStatus = "watched"
	StatusEpQueue   EpStatus = "queue"
	StatusEpDrop    EpStatus = "drop"
	StatusEpRemove  EpStatus = "remove"
)

// ResponseGroup ...
type ResponseGroup string

// ...
const (
	ResponseSmall  ResponseGroup = "small"
	ResponseMediam ResponseGroup = "mediam"
	ResponseLarge  ResponseGroup = "large"
)

// StatusCode ...
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
