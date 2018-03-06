package chii

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	httpClient := &http.Client{}
	c := NewClient(httpClient)
	uc := newUserService(c.sling)
	user, _, err := uc.Get("everpcpc")
	assert.Nil(t, err)
	assert.Equal(t, 26024, user.ID)
}
