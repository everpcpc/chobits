package chii

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEpisode(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{})
	ep, _, err := c.Episode.Get(1088292)
	r.Nil(err)
	r.Equal(1088292, ep.ID)
	r.Equal(343656, ep.SubjectID)
	r.Equal("近すぎじゃね？", ep.Name)
}
