package chii

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/Collection-API.md

// CollectionService provides methods for accessing Bangumi collection API endpoints.
type CollectionService struct {
	sling *sling.Sling
}

func newCollectionService(sling *sling.Sling) *CollectionService {
	return &CollectionService{
		sling: sling.Path("collection/"),
	}
}

// Collection ...
type Collection struct {
	Status    UserCollectStatus `json:"status,omitempty"`
	Rating    int               `json:"rating,omitempty"`
	Comment   string            `json:"comment,omitempty"`
	Private   int               `json:"private,omitempty"`
	Tags      []string          `json:"tags,omitempty"`
	EpStatus  int               `json:"ep_status,omitempty"`
	Lasttouch int               `json:"lasttouch,omitempty"`
	User      User              `json:"user,omitempty"`
}

///////////////////////////////////////////////////////////////

// Info returns the Collection status for a requests subject.
func (s *CollectionService) Info(subjectID int) (Collection, *http.Response, error) {
	collection := Collection{}
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(subjectID)).Receive(&collection, apiError)
	return collection, resp, relevantError(err, *apiError)
}

///////////////////////////////////////////////////////////////

type collectionStatusBody struct {
	Status  CollectionsStatus `url:"status"`
	Comment string            `url:"comment,omitempty"`
	Tags    []string          `url:"tags,omitempty,comma"`
	Rating  int               `url:"rating,omitempty"`
	Privacy int               `url:"privacy,omitempty"`
}

// Manage updates a collection status.
func (s *CollectionService) Manage(subjectID int, status CollectionsStatus, comment string, tags []string, rating int, private bool) (Collection, *http.Response, error) {
	apiError := new(APIError)
	collection := Collection{}
	path := fmt.Sprintf("%d/update", subjectID)
	body := &collectionStatusBody{
		Status:  status,
		Comment: comment,
		Tags:    tags,
		Rating:  rating,
	}
	if private {
		body.Privacy = 1
	}
	resp, err := s.sling.New().Post(path).BodyForm(body).Receive(&collection, apiError)
	return collection, resp, relevantError(err, *apiError)
}
