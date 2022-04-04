package chii

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

// https://bangumi.github.io/api/#/%E6%9D%A1%E7%9B%AE

type SubjectService struct {
	sling *sling.Sling
}

func newSubjectService(sling *sling.Sling) *SubjectService {
	return &SubjectService{
		sling: sling.Path("subjects/"),
	}
}

type Subject struct {
	ID            int               `json:"id"`
	Type          SubjectType       `json:"type"`
	Name          string            `json:"name"`
	NameCN        string            `json:"name_cn"`
	Summary       string            `json:"summary"`
	NSFW          bool              `json:"nsfw"`
	Locked        bool              `json:"locked"`
	Date          string            `json:"date,omitempty"`
	Platform      string            `json:"platform"`
	Images        Images            `json:"images,omitempty"`
	Infobox       []Infobox         `json:"infobox,omitempty"`
	Volumes       int               `json:"volumes"`
	Eps           int               `json:"eps"`
	TotalEpisodes int               `json:"total_episodes"`
	Rating        SubjectRating     `json:"rating"`
	Collection    SubjectCollection `json:"collection"`
	Tags          []SubjectTag      `json:"tags"`
}

type Infobox struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type SubjectRating struct {
	Rank  int                `json:"rank"`
	Total int                `json:"total"`
	Count SubjectRatingCount `json:"count"`
	Score float64            `json:"score"`
}

type SubjectRatingCount struct {
	One   int `json:"1"`
	Two   int `json:"2"`
	Three int `json:"3"`
	Four  int `json:"4"`
	Five  int `json:"5"`
	Six   int `json:"6"`
	Seven int `json:"7"`
	Eight int `json:"8"`
	Nine  int `json:"9"`
	Ten   int `json:"10"`
}

type SubjectCollection struct {
	Wish    int `json:"wish"`
	Collect int `json:"collect"`
	Doing   int `json:"doing"`
	OnHold  int `json:"on_hold"`
	Dropped int `json:"dropped"`
}

type SubjectTag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

///////////////////////////////////////////////////////////////

type subjectParams struct {
	ResponseGroup ResponseGroup `url:"responseGroup,omitempty"`
}

func (s *SubjectService) Info(id int) (Subject, *http.Response, error) {
	subject := Subject{}
	apiError := APIError{}
	resp, err := s.sling.New().Get(strconv.Itoa(id)).Receive(&subject, &apiError)
	return subject, resp, relevantError(err, apiError)
}
