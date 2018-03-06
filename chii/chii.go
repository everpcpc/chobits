package chii

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a Bangumi client for making Bangumi API requests.
type Client struct {
	sling *sling.Sling
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(bgmAPI)
	return &Client{
		sling: base,
	}
}
