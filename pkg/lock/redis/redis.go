package redis

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"github.com/saromanov/locker/pkg/lock"
)

const (
	lockerKey = "locker"
)

// redisLock provides implementation of locks from redis
type redisLock struct {
	client *redis.Client
}

func New(conf lock.Config) lock.Locker {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: "",
		DB:       0,
	})
	return &redisLock{
		client: client,
	}
}

func (t *redisLock) Lock() bool {
	resp := t.client.SetNX(lockerKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil {
		fmt.Println("unable to apply lock: ", err)
		return false
	}
	if !lockSuccess {
		fmt.Println("unable to apply lock")
		return false
	}

	return true
}

func (t *redisLock) Unlock() {
	delResp := t.client.Del(lockerKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}
