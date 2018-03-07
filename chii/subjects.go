package chii

import (
	"github.com/dghubble/sling"
)

// https://github.com/bangumi/api/blob/master/docs-raw/Subject-API.md

// SubjectService provides methods for accessing Bangumi subject API endpoints.
type SubjectService struct {
	sling *sling.Sling
}

// TODO:(everpcpc)

// SubjectOverview represents a Bangumi Subject.
type SubjectOverview struct {
	SubjectID int     `json:"subject_id"`
	Eps       []Ep    `json:"eps,omitempty"`
	Detail    Subject `json:"subject,omitempty"`
}

// Subject respents a Bangumi subject detail
type Subject struct {
	ID          int               `json:"id"`
	URL         string            `json:"url"`
	Type        SubjectType       `json:"type"`
	Name        string            `json:"name"`
	NameCN      string            `json:"name_cn,omitempty"`
	Summary     string            `json:"summary,omitempty"`
	Eps         int               `json:"eps,omitempty"`
	EpsCount    int               `json:"eps_count,omitempty"`
	AirDate     string            `json:"air_date,omitempty"`
	AirWeekdays int               `json:"air_weekday,omitempty"`
	Images      SubjectImages     `json:"images,omitempty"`
	Collection  SubjectCollection `json:"collection,omitempty"`
}

// SubjectCollection represents collection for a subject
type SubjectCollection struct {
	Doing int `json:"doing,omitempty"`
}

// SubjectImages respents a set of Bangumi subject images
type SubjectImages struct {
	Large  string `json:"large,omitempty"`
	Common string `json:"common,omitempty"`
	Medium string `json:"medium,omitempty"`
	Small  string `json:"small,omitempty"`
	Grid   string `json:"grid,omitempty"`
}

// Ep represents a Bangumi Ep.
type Ep struct {
	ID     int           `json:"id"`
	Status EpWatchStatus `json:"status"`
}

// EpWatchStatus represents a Bangumi Ep status.
type EpWatchStatus struct {
	ID      int    `json:"id"`
	CSSName string `json:"css_name"`
	URLName string `json:"url_name"`
	CNName  string `json:"cn_name"`
}
