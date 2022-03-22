package cache

import (
	"go.uber.org/fx"
	iface "red-packet/util/interface"
	"red-packet/util/redis"
)

type Params struct {
	fx.In

	Cache redis.Redis
}

var Module = fx.Options(
	fx.Provide(
		NewCache,
	),
)

func NewCache(p Params) iface.ICache {
	return &redis.Cache{
		Redis: p.Cache,
	}
}
