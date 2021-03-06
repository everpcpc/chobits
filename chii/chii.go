package chii

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a Bangumi client for making Bangumi API requests.
type Client struct {
	sling      *sling.Sling
	clientID   string
	token      string
	Users      *UserService
	Subject    *SubjectService
	Collection *CollectionService
}

type appIDParam struct {
	AppID string `url:"app_id,omitempty"`
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, clientID, token string) *Client {
	params := &appIDParam{AppID: clientID}
	base := sling.New().Client(httpClient).Base(bgmAPI).Set("User-Agent", "Go Chii Client").Set("Authorization", "Bearer "+token).QueryStruct(params)
	return &Client{
		sling:      base,
		clientID:   clientID,
		token:      token,
		Users:      newUserService(base.New()),
		Subject:    newSubjectService(base.New()),
		Collection: newCollectionService(base.New()),
	}
}

// Images respents a set of Bangumi images
type Images struct {
	Large  string `json:"large,omitempty"`
	Common string `json:"common,omitempty"`
	Medium string `json:"medium,omitempty"`
	Small  string `json:"small,omitempty"`
	Grid   string `json:"grid,omitempty"`
}

// Info represents a person/character info
type Info struct {
	Birth  string `json:"birth,omitempty"`
	Height string `json:"height,omitempty"`
	Gender string `json:"gender,omitempty"`
	Alias  Alias  `json:"alias,omitempty"`

	// FIXME: some times string https://github.com/bangumi/api/issues/16
	Source  []string `json:"source,omitempty"`
	NameCN  string   `json:"name_cn,omitempty"`
	CV      string   `json:"cv,omitempty"`
	Twitter string   `json:"twitter,omitempty"`
}

// Alias ...
type Alias struct {
	JP     string `json:"jp,omitempty"`
	Romaji string `json:"romaji,omitempty"`
	ZH     string `json:"zh,omitempty"`
	Kana   string `json:"kana,omitempty"`
}

// TODO: 每日放送
// GET /calendar
