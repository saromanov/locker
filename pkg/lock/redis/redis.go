package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/saromanov/locker/pkg/lock"
	log "github.com/sirupsen/logrus"
)

const (
	lockerKey = "locker"
)

// redisLock provides implementation of locks from redis
type redisLock struct {
	client *redis.Client
}

func New(conf *lock.Config) lock.Locker {
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
		log.WithError(err).WithField("lock", "redis").Error("unable to apply lock")
		return false
	}
	if !lockSuccess {
		log.WithField("lock", "redis").Error("unable to apply lock")
		return false
	}

	return true
}

func (t *redisLock) Unlock() {
	delResp := t.client.Del(lockerKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		log.WithField("lock", "redis").Info("lock is successed")
	} else {
		log.WithField("lock", "redis").Error("unable to apply lock")
	}
}
