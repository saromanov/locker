package trivial

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	l := New()
	assert.Equal(t, l.Lock(), true)
}

func TestLocks(t *testing.T) {
	var l = New()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				assert.Fail(t, "lock is failed")
				return
			}
			l.Unlock()
		}()
	}
	wg.Wait()
}
