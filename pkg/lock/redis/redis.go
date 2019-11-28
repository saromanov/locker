package redis

import (
	"github.com/go-redis/redis"
	"github.com/saromanov/locker/pkg/lock"
)

// redisLock provides implementation of locks from redis
type redisLock struct {
}

func New() *lock.Locker {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &redisLock{}
}
