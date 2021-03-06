package chii

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubjectInfo(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)
	subject, _, err := c.Subject.Info(12, ResponseLarge)
	r.Nil(err)
	r.Equal(12, subject.ID)
}
