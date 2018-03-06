package chii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubjectType(t *testing.T) {
	assert.Equal(t, SubjectType(1), TypeBook)
	assert.Equal(t, SubjectType(3), TypeMusic)
	assert.Equal(t, SubjectType(6), TypeReal)
}

func TestCollectionStatus(t *testing.T) {
	assert.Equal(t, CollectionsStatus("dropped"), StatusDropped)
	assert.Equal(t, CollectionsStatus("on_hold"), StatusOnHold)
}
