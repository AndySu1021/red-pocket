package lock

import (
	"github.com/bsm/redislock"
	"red-packet/util/redis"
)

func NewRedisLocker(redis redis.Redis) *redislock.Client {
	return redislock.New(redis)
}
