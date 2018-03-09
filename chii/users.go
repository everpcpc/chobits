package chii

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/User-API.md

// User represents a Bangumi User.
type User struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   Images `json:"avatar"`
	Sign     string `json:"sign"`
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

// SubjectOverview ...
type SubjectOverview struct {
	SubjectID int     `json:"subject_id"`
	Eps       []Ep    `json:"eps,omitempty"`
	Detail    Subject `json:"subject,omitempty"`
}

// UserCollectionsStatus ...
type UserCollectionsStatus struct {
	Type     SubjectType   `json:"type,omitempty"`
	Name     string        `json:"name,omitempty"`
	NameCN   string        `json:"name_cn,omitempty"`
	Collects []UserCollect `json:"collects,omitempty"`
}

// UserCollect ...
type UserCollect struct {
	Status UserCollectStatus `json:"status,omitempty"`
	Count  int               `json:"count,omitempty"`
	List   []SubjectOverview `json:"list,omitempty"`
}

// UserCollectStatus ...
type UserCollectStatus struct {
	ID   int               `json:"id,,omitempty"`
	Type CollectionsStatus `json:"type,omitempty"`
	Name string            `json:"name,omitempty"`
}

// UserCollectionsSubject ...
type UserCollectionsSubject struct {
	Name      string  `json:"name,omitempty"`
	SubjectID int     `json:"subject_id,omitempty"`
	EpStatus  int     `json:"ep_status,omitempty"`
	VolStatus int     `json:"vol_status,omitempty"`
	LastTouch int     `json:"lasttouch,omitempty"`
	Subject   Subject `json:"subject,omitempty"`
}

///////////////////////////////////////////////////////////////

// Info returns the requested User.
func (s *UserService) Info(username string) (User, *http.Response, error) {
	user := User{}
	apiError := new(APIError)
	resp, err := s.sling.New().Get(username).Receive(&user, apiError)
	return user, resp, relevantError(err, *apiError)
}

///////////////////////////////////////////////////////////////

type userCollectionParams struct {
	cat           CollectionType `url:"cat,omitempty"`
	ids           string         `url:"ids,omitempty"`
	responseGroup ResponseGroup  `url:"responseGroup,omitempty"`
}

// Collection returns the collection for requested User.
// @username
// @cat CollectionType
// @ids: exp 1,2,4,6
// @responseGroup: medium/small, default medium
func (s *UserService) Collection(username string, cat CollectionType, ids []int, responseGroup ResponseGroup) ([]UserCollectionsSubject, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/collection", username)
	params := &userCollectionParams{
		cat:           cat,
		ids:           strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"),
		responseGroup: responseGroup,
	}
	subjects := []UserCollectionsSubject{}
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(&subjects, apiError)
	return subjects, resp, relevantError(err, *apiError)
}

///////////////////////////////////////////////////////////////

type userCollectionsOverviewParams struct {
	maxResults int `url:"max_results,omitempty"`
}

// CollectionsOverview returns the collection overview
// for requested User, not paged.
// @max_results: max 25
func (s *UserService) CollectionsOverview(username string, subjectType SubjectType, maxResults int) ([]UserCollectionsStatus, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/collections/%s", username, subjectTypeString[subjectType])
	params := &userCollectionsOverviewParams{maxResults: maxResults}
	status := []UserCollectionsStatus{}
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(&status, apiError)
	return status, resp, relevantError(err, *apiError)
}

///////////////////////////////////////////////////////////////

// CollectionsStatus returns all collections status
// for the requested User.
func (s *UserService) CollectionsStatus(username string) ([]UserCollectionsStatus, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/collections/status", username)
	status := []UserCollectionsStatus{}
	resp, err := s.sling.New().Get(path).Receive(&status, apiError)
	return status, resp, relevantError(err, *apiError)
}

///////////////////////////////////////////////////////////////

type userProgressParams struct {
	subjectID int `url:"subject_id,omitempty"`
}

// Progress returns progress for the requested User.
// limit 15
func (s *UserService) Progress(username string) ([]SubjectOverview, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/progress", username)

	subjects := []SubjectOverview{}
	resp, err := s.sling.New().Get(path).Receive(&subjects, apiError)
	return subjects, resp, relevantError(err, *apiError)
}

// ProgressSubject returns progress for the requested User
// and for a given subject.
func (s *UserService) ProgressSubject(username string, subjectID int) (SubjectOverview, *http.Response, error) {
	apiError := new(APIError)
	path := fmt.Sprintf("%s/progress", username)
	subject := SubjectOverview{}

	params := &userProgressParams{subjectID: subjectID}
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(&subject, apiError)
	return subject, resp, relevantError(err, *apiError)
}
