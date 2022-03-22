package redis

import (
	"context"
	"time"
)

type Cache struct {
	Redis Redis
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	result, err := c.Redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *Cache) SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (exists bool, err error) {
	ok, err := c.Redis.SetNX(ctx, key, data, expireAt).Result()
	if err != nil {
		return false, err
	}
	return ok, err
}

func (c *Cache) SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	if err := c.Redis.SetEX(ctx, key, data, expireAt).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cache) LPush(ctx context.Context, key string, values ...interface{}) (total int64, err error) {
	total, err = c.Redis.LPush(ctx, key, values).Result()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (c *Cache) RPop(ctx context.Context, key string) (result string, err error) {
	result, err = c.Redis.RPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *Cache) Expire(ctx context.Context, key string, expireAt time.Duration) error {
	err := c.Redis.Expire(ctx, key, expireAt).Err()
	if err != nil {
		return err
	}

	return nil
}
