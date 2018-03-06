package chii

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	c := NewClient(&http.Client{})
	user, _, err := c.Users.Info("everpcpc")
	assert.Nil(t, err)
	assert.Equal(t, 26024, user.ID)
}
