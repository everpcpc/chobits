package chii

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a Bangumi client for making Bangumi API requests.
type Client struct {
	sling    *sling.Sling
	clientID string
	token    string
	Users    *UserService
}

type appIDParam struct {
	AppID string `url:"app_id,omitempty"`
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, clientID, token string) *Client {
	params := &appIDParam{AppID: clientID}
	base := sling.New().Client(httpClient).Base(bgmAPI).Set("User-Agent", "Go Chii Client").Set("Authorization", "Bearer "+token).QueryStruct(params)
	return &Client{
		sling:    base,
		clientID: clientID,
		token:    token,
		Users:    newUserService(base.New()),
	}
}
