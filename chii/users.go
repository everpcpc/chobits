package chii

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/User-API.md

// User represents a Bangumi User.
type User struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   Avatar `json:"avatar"`
	Sign     string `json:"sign"`
}

// Avatar represents avatars for a Bangumi User.
type Avatar struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Small  string `json:"small"`
}

// UserService provides methods for accessing Bangumi user API endpoints.
type UserService struct {
	sling *sling.Sling
}

func newUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path("user/"),
	}
}

// Info returns the requested User.
func (s *UserService) Info(username string) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(username).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}

// Collection returns the collection for requested User.
// TODO:(everpcpc)
// @username
// @cat CollectionType
// @ids: exp 1,2,4,6
// @responseGroup: medium/small, default medium
func (s *UserService) Collection(username string, cat CollectionType, ids []int, responceGroup ResponseGroup) (*User, *http.Response, error) {
	return nil, nil, nil
}

// CollectionsOverview returns the collection overview for requested User, not paged.
// TODO:(everpcpc) Error: APP ID Parameter is Missing"
// @username
// @subject_type
// @max_result: max 25
// @responseGroup: medium/small, default medium
func (s *UserService) CollectionsOverview(username string, subjectType SubjectType, maxResult int) (interface{}, *http.Response, error) {
	return nil, nil, nil
}

// CollectionsStatus returns all collections status for the requested User.
// TODO:(everpcpc) Error: APP ID Parameter is Missing
func (s *UserService) CollectionsStatus(username string) (interface{}, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/collections/status", username)
	resp, err := s.sling.New().Get(path).Receive(nil, apiError)
	return nil, resp, relevantError(err, *apiError)
}

// Progress returns progress for the requested User.
// TODO:(everpcpc) Requires Authentication
// @username
// @subject_id
func (s *UserService) Progress(username string, subjectID int) (interface{}, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/progress", username)
	resp, err := s.sling.New().Get(path).Receive(nil, apiError)
	return nil, resp, relevantError(err, *apiError)
}
