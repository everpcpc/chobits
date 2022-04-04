package chii

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

type EpisodeType int

const (
	EpisodeTypeBase EpisodeType = iota
	EpisodeTypeSP
	EpisodeTypeOP
	EpisodeTypeED
)

type EpisodeStatus string

const (
	EpisodeStatusEpWatched EpisodeStatus = "watched"
	EpisodeStatusEpQueue   EpisodeStatus = "queue"
	EpisodeStatusEpDrop    EpisodeStatus = "drop"
	EpisodeStatusEpRemove  EpisodeStatus = "remove"
)

type EpisodeService struct {
	sling *sling.Sling
}

func newEpisodeService(sling *sling.Sling) *EpisodeService {
	return &EpisodeService{
		sling: sling.Path("episodes/"),
	}
}

type EpisodeDetail struct {
	ID int `json:"id"`

	/// 0 本篇，1 SP，2 OP，3 ED
	Type EpisodeType `json:"type"`

	Name   string `json:"name"`
	NameCN string `json:"name_cn"`

	/// 同类条目的排序和集数
	Sort int `json:"sort"`

	/// 条目内的集数, 从1开始。非本篇剧集的此字段无意义
	Ep int `json:"ep,omitempty"`

	Airdate  string `json:"airdate"`
	Comment  int    `json:"comment"`
	Duration string `json:"duration"`

	/// 简介
	Desc string `json:"desc"`

	/// 音乐曲目的碟片数
	Disc int `json:"disc"`

	SubjectID int `json:"subject_id"`
}

///////////////////////////////////////////////////////////////

func (s *EpisodeService) Get(id int) (EpisodeDetail, *http.Response, error) {
	episode := EpisodeDetail{}
	apiError := APIError{}
	resp, err := s.sling.New().Get(strconv.Itoa(id)).Receive(&episode, &apiError)
	return episode, resp, relevantError(err, apiError)
}
