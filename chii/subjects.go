package chii

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/Subject-API.md

// SubjectService provides methods for accessing Bangumi subject API endpoints.
type SubjectService struct {
	sling *sling.Sling
}

func newSubjectService(sling *sling.Sling) *SubjectService {
	return &SubjectService{
		sling: sling.Path("subject/"),
	}
}

// Subject respents a Bangumi subject detail
type Subject struct {
	ID          int               `json:"id"`
	URL         string            `json:"url"`
	Type        SubjectType       `json:"type"`
	Name        string            `json:"name"`
	NameCN      string            `json:"name_cn,omitempty"`
	Summary     string            `json:"summary,omitempty"`
	Eps         []Ep              `json:"eps,omitempty"`
	EpsCount    int               `json:"eps_count,omitempty"`
	AirDate     string            `json:"air_date,omitempty"`
	AirWeekdays int               `json:"air_weekday,omitempty"`
	Rating      SubjectRating     `json:"rating,omitempty"`
	Rank        int               `json:"rank,omitempty"`
	Images      Images            `json:"images,omitempty"`
	Collection  SubjectCollection `json:"collection,omitempty"`
	Character   []Character       `json:"crt,omitempty"`
	Staff       []Person          `json:"person,omitempty"`
	Topic       []Topic           `json:"topic,omitempty"`
	Blog        []Blog            `json:"blog,omitempty"`
}

// SubjectRating ...
type SubjectRating struct {
	Total int                `json:"total,omitempty"`
	Count SubjectRatingCount `json:"count,omitempty"`
	Score float64            `json:"score,omitempty"`
}

// SubjectRatingCount ...
type SubjectRatingCount struct {
	One   int `json:"1,omitempty"`
	Two   int `json:"2,omitempty"`
	Three int `json:"3,omitempty"`
	Four  int `json:"4,omitempty"`
	Five  int `json:"5,omitempty"`
	Six   int `json:"6,omitempty"`
	Seven int `json:"7,omitempty"`
	Eight int `json:"8,omitempty"`
	Nine  int `json:"9,omitempty"`
	Ten   int `json:"10,omitempty"`
}

// SubjectCollection represents collection for a subject
type SubjectCollection struct {
	Wish    int `json:"wish,omitempty"`
	Collect int `json:"collect,omitempty"`
	Doing   int `json:"doing,omitempty"`
	OnHold  int `json:"on_hold,omitempty"`
	Dropped int `json:"dropped,omitempty"`
}

// Ep represents a Bangumi Ep.
type Ep struct {
	ID     int           `json:"id"`
	Status EpWatchStatus `json:"status,omitempty"`
}

// EpWatchStatus represents a Bangumi Ep status.
type EpWatchStatus struct {
	ID      int      `json:"id"`
	CSSName string   `json:"css_name,omitempty"`
	URLName string   `json:"url_name,omitempty"`
	CNName  string   `json:"cn_name,omitempty"`
	Status  EpStatus `json:"status,omitempty"`
}

///////////////////////////////////////////////////////////////

// Info returns the requested Subject.
func (s *SubjectService) Info(id int, responseGroup ResponseGroup) (Subject, *http.Response, error) {
	subject := Subject{}
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(id)).Receive(&subject, apiError)
	return subject, resp, relevantError(err, *apiError)
}

///////////////////////////////////////////////////////////////

// InfoEp returns the Eps and basic info for requested Subject.
func (s *SubjectService) InfoEp(id int, responseGroup ResponseGroup) (Subject, *http.Response, error) {
	subject := Subject{}
	apiError := new(APIError)
	resp, err := s.sling.New().Get(strconv.Itoa(id)+"/ep").Receive(&subject, apiError)
	return subject, resp, relevantError(err, *apiError)
}
