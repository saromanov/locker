package redis

import (
	"testing"

	"github.com/saromanov/locker/pkg/lock"
	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	l := New(lock.Config{})
	assert.Equal(t, l.Lock(), true)
}
