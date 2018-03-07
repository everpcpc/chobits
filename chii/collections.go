package chii

import (
	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/Collection-API.md

// CollectionService provides methods for accessing Bangumi collection API endpoints.
type CollectionService struct {
	sling *sling.Sling
}

// TODO:(everpcpc)
