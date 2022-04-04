package chii

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSubject(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{})
	subject, _, err := c.Subject.Get(343656)
	r.Nil(err)
	r.Equal(343656, subject.ID)
	r.Equal("阿波連さんははかれない", subject.Name)
}
