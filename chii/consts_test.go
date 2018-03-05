package chii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubjectType(t *testing.T) {
	assert.Equal(t, TypeBook, SubjectType(1))
	assert.Equal(t, TypeMusic, SubjectType(3))
	assert.Equal(t, TypeReal, SubjectType(6))
}

func TestCollectionStatus(t *testing.T) {
	assert.Equal(t, StatusDropped, CollectionsStatus("dropped"))
	assert.Equal(t, StatusOnHold, CollectionsStatus("on_hold"))
}
