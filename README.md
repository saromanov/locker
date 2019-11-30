# locker

Implementation of some types of locks

## Usage

```go
package main

import (
	"fmt"
	"sync"

	"github.com/saromanov/locker/pkg/lock"
	"github.com/saromanov/locker/pkg/lock/redis"
	"github.com/saromanov/locker/pkg/lock/trivial"
)

func trivialLock() {
	var counter int
	var l = trivial.New()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}

func redisLock() {
	re := redis.New(&lock.Config{
		Address: "localhost:6379",
	})
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(re.Lock())
			re.Unlock()
		}()
	}
	wg.Wait()
}
func main() {
	redisLock()
}
```