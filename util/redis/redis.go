package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

// Redis 提供操作 redis 的介面
type Redis interface {
	redis.Cmdable
	GetClient() *redis.Client
	GetConfig() *Config
}

type Client struct {
	*redis.Client
	Config *Config
}

func (c *Client) GetClient() *redis.Client {
	return c.Client
}

func (c *Client) GetConfig() *Config {
	return c.Config
}

func NewRedisClient(cfg *Config) (Redis, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	if len(cfg.Addresses) == 0 {
		return nil, fmt.Errorf("redis config address is empty")
	}

	var client *redis.Client
	err := backoff.Retry(func() error {
		client = redis.NewClient(&redis.Options{
			Addr:       cfg.Addresses[0],
			Password:   cfg.Password,
			MaxRetries: cfg.MaxRetries,
			PoolSize:   cfg.PoolSizePerNode,
			DB:         cfg.DB,
		})
		err := client.Ping(context.Background()).Err()
		if err != nil {
			log.Error().Msgf("ping occurs error after connecting to redis: %s", err)
			return fmt.Errorf("ping occurs error after connecting to redis: %s", err)
		}
		return nil
	}, bo)

	if err != nil {
		return nil, err
	}

	return &Client{Client: client, Config: cfg}, nil
}
