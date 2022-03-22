package iface

import (
	"context"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (exists bool, err error)
	SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error
	LPush(ctx context.Context, key string, values ...interface{}) (total int64, err error)
	RPop(ctx context.Context, key string) (result string, err error)
	Expire(ctx context.Context, key string, expireAt time.Duration) error
}
