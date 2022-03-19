package red_packet

import (
	"github.com/bsm/redislock"
	"go.uber.org/fx"
	iface "red-packet/pkg/interface"
)

type service struct {
	repo      iface.IRepository
	cache     iface.ICache
	redisLock *redislock.Client
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type Params struct {
	fx.In

	Repo      iface.IRepository
	Cache     iface.ICache
	RedisLock *redislock.Client
}

func New(p Params) iface.IRedPacketService {
	return &service{
		repo:      p.Repo,
		cache:     p.Cache,
		redisLock: p.RedisLock,
	}
}
