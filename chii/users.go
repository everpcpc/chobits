package chii

import (
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

// newUserService returns a new UserService.
func newUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path("user/"),
	}
}

// Get returns the requested User.
func (s *UserService) Get(username string) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(username).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}
