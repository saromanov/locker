package trivial

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestLock(t *testing.T) {
	l := New()
	assert.Equal(t, l.Lock(), true)
}