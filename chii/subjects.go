package chii

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

type SubjectType int

const (
	SubjectTypeNil SubjectType = iota
	SubjectTypeBook
	SubjectTypeAnime
	SubjectTypeMusic
	SubjectTypeGame
	SubjectTypeReserved
	SubjectTypeReal
)

var subjectTypeString = map[SubjectType]string{
	SubjectTypeNil:      "nil",
	SubjectTypeBook:     "book",
	SubjectTypeAnime:    "anime",
	SubjectTypeMusic:    "music",
	SubjectTypeGame:     "game",
	SubjectTypeReserved: "reserved",
	SubjectTypeReal:     "real",
}

type SubjectService struct {
	sling *sling.Sling
}

func newSubjectService(sling *sling.Sling) *SubjectService {
	return &SubjectService{
		sling: sling.Path("subjects/"),
	}
}

type Subject struct {
	ID      int         `json:"id"`
	Type    SubjectType `json:"type"`
	Name    string      `json:"name"`
	NameCN  string      `json:"name_cn"`
	Summary string      `json:"summary"`
	NSFW    bool        `json:"nsfw"`
	Locked  bool        `json:"locked"`

	/// air date in YYYY-MM-DD format
	Date string `json:"date,omitempty"`

	/// TV, Web, 欧美剧, PS4...
	Platform string `json:"platform"`

	Images  Images    `json:"images,omitempty"`
	Infobox []Infobox `json:"infobox,omitempty"`

	/// 书籍条目的册数，由旧服务端从wiki中解析
	Volumes int `json:"volumes"`

	/// 由旧服务端从wiki中解析，对于书籍条目为话数
	Eps int `json:"eps"`

	/// 数据库中的章节数量
	TotalEpisodes int `json:"total_episodes"`

	Rating     Rating     `json:"rating"`
	Collection Collection `json:"collection"`
	Tags       []Tag      `json:"tags"`
}

type Infobox struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Rating struct {
	Rank  int         `json:"rank"`
	Total int         `json:"total"`
	Count RatingCount `json:"count"`
	Score float64     `json:"score"`
}

type RatingCount struct {
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

type Collection struct {
	Wish    int `json:"wish"`
	Collect int `json:"collect"`
	Doing   int `json:"doing"`
	OnHold  int `json:"on_hold"`
	Dropped int `json:"dropped"`
}

type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

///////////////////////////////////////////////////////////////

func (s *SubjectService) Get(id int) (Subject, *http.Response, error) {
	subject := Subject{}
	apiError := APIError{}
	resp, err := s.sling.New().Get(strconv.Itoa(id)).Receive(&subject, &apiError)
	return subject, resp, relevantError(err, apiError)
}
