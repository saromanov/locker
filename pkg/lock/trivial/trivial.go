package trivial

import (
	"github.com/saromanov/locker/pkg/lock"
)

type trivial struct {
	c chan struct{}
}

// New generate a try lock
func New() lock.Locker {
	c := make(chan struct{}, 1)
	c <- struct{}{}
	return &trivial {
		c: c,
	}
}

// Lock provides locking of the data
func (t *trivial) Lock() bool {
	lockResult := false
	select {
	case <-t.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock provides unlocking of the data
func (t *trivial) Unlock() {
	t.c <- struct{}{}
}
