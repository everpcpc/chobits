package chii

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/Progress-API.md

// ProgressService provides methods for accessing Bangumi progress API endpoints.
type ProgressService struct {
	sling *sling.Sling
}

func newProgressService(sling *sling.Sling) *ProgressService {
	return &ProgressService{
		sling: sling.Path(""),
	}
}

// Update update a ep status.
// * Requires Authentication
// TODO:(everpcpc)
func (s *ProgressService) Update(epID int, status EpStatus) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}

// UpdateTo update a subject progress to provided ep.
// * Requires Authentication
// TODO:(everpcpc)
func (s *ProgressService) UpdateTo(subjectID int, watchedEps string, watchedVols string) (*http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("subject/%d/update/watched_eps", subjectID)
	resp, err := s.sling.New().Post(path).Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
