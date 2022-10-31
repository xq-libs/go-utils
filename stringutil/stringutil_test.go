package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var a string
  assert.Equal(t, true, IsEmpty(a), "The nil str should be empty")
	assert.Equal(t, true, IsEmpty(""), "The empty str should be empty")
	assert.Equal(t, false, IsEmpty(" "), "The blank str should not be empty")
	assert.Equal(t, false, IsEmpty("abc"), "The abc str be should not be empty")
}