package chii

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Client struct {
	sling    *sling.Sling
	clientID string
	token    string
	Subject  *SubjectService
	Episode  *EpisodeService
}

type appIDParam struct {
	AppID string `url:"app_id,omitempty"`
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(bgmAPI).Set("User-Agent", "Chii-1.0")
	return &Client{
		sling:   base,
		Subject: newSubjectService(base.New()),
		Episode: newEpisodeService(base.New()),
	}
}

func NewClientWithAuth(httpClient *http.Client, clientID, token string) *Client {
	params := &appIDParam{AppID: clientID}
	base := sling.New().Client(httpClient).Base(bgmAPI).Set("User-Agent", "Chii-1.0").Set("Authorization", "Bearer "+token).QueryStruct(params)
	return &Client{
		sling:    base,
		clientID: clientID,
		token:    token,
		Subject:  newSubjectService(base.New()),
		Episode:  newEpisodeService(base.New()),
	}
}

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

type Alias struct {
	JP     string `json:"jp,omitempty"`
	Romaji string `json:"romaji,omitempty"`
	ZH     string `json:"zh,omitempty"`
	Kana   string `json:"kana,omitempty"`
}
