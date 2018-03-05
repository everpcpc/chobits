package chii

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
	StatusWatched EpStatus = "watched"
	StatusQueue   EpStatus = "queue"
	StatusDrop    EpStatus = "drop"
	StatusRemove  EpStatus = "remove"
)

// ResponseGroup ...
type ResponseGroup string

// ...
const (
	ResponceSmall  ResponseGroup = "small"
	ResponceMediam ResponseGroup = "mediam"
	ResponceLarge  ResponseGroup = "large"
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
