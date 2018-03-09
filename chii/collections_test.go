package chii

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollectionInfo(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)
	collection, _, err := c.Collection.Info(12)
	r.Nil(err)
	r.Equal(10, collection.Rating)
	r.Equal(StatusCollect, collection.Status.Type)
}

func TestCollectionManage(t *testing.T) {
	r := require.New(t)
	c := NewClient(&http.Client{}, testClientID, testToken)
	collection, _, err := c.Collection.Manage(12, StatusCollect, "", []string{"CLAMP", "科幻", "治愈"}, 10, false)
	r.Nil(err)
	r.Nil(nil, collection.Status.Type)
}
