package chii

import (
	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/Search-API.md

// SearchService provides methods for accessing Bangumi search API endpoints.
type SearchService struct {
	sling *sling.Sling
}

// TODO:(everpcpc)
