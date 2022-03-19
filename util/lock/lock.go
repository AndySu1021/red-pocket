package lock

import (
	"demo/util/redis"
	"github.com/bsm/redislock"
)

func NewRedisLocker(redis redis.Redis) *redislock.Client {
	return redislock.New(redis)
}
