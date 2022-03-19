package cache

import (
	"context"
	iface "demo/pkg/interface"
	"demo/util/redis"
	"go.uber.org/fx"
	"time"
)

type cache struct {
	redis redis.Redis
}

func (c *cache) Get(ctx context.Context, key string) (string, error) {
	result, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *cache) SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (exists bool, err error) {
	ok, err := c.redis.SetNX(ctx, key, data, expireAt).Result()
	if err != nil {
		return false, err
	}
	return ok, err
}

func (c *cache) SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	if err := c.redis.SetEX(ctx, key, data, expireAt).Err(); err != nil {
		return err
	}
	return nil
}

func (c *cache) LPush(ctx context.Context, key string, values ...interface{}) (total int64, err error) {
	total, err = c.redis.LPush(ctx, key, values).Result()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (c *cache) RPop(ctx context.Context, key string) (result string, err error) {
	result, err = c.redis.RPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *cache) Expire(ctx context.Context, key string, expireAt time.Duration) error {
	err := c.redis.Expire(ctx, key, expireAt).Err()
	if err != nil {
		return err
	}

	return nil
}

type Params struct {
	fx.In

	Redis redis.Redis
}

var Module = fx.Options(
	fx.Provide(
		NewRedis,
	),
)

func NewRedis(p Params) iface.ICache {
	return &cache{
		redis: p.Redis,
	}
}
