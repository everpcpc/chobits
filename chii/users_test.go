package chii

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserInfo(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)
	user, _, err := c.Users.Info("everpcpc")
	r.Nil(err)
	r.Equal(26024, user.ID)
}

func TestUserCollection(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)
	subjects, _, err := c.Users.Collection("everpcpc", TypeWatching, nil, ResponseSmall)
	r.Nil(err)
	r.Equal(50, len(subjects))
}

func TestUserCollectionsOverview(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)
	statuses, _, err := c.Users.CollectionsOverview("everpcpc", TypeAnime, 1)
	r.Nil(err)
	r.NotEmpty(statuses)
	r.Equal(TypeAnime, statuses[0].Type)
	r.NotEmpty(statuses[0].Collects)
}

func TestUserCollectionsStatus(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)
	statuses, _, err := c.Users.CollectionsStatus("everpcpc")
	r.Nil(err)
	r.NotEmpty(statuses)
}

func TestUserProgress(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)

	subjects, _, err := c.Users.Progress("everpcpc")
	r.Nil(err)
	r.Equal(15, len(subjects))

	subject, _, err := c.Users.ProgressSubject("everpcpc", 12)
	r.Nil(err)
	r.Equal(12, subject.SubjectID)
}
