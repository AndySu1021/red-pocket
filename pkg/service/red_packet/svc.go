package red_packet

import (
	"github.com/bsm/redislock"
	"go.uber.org/fx"
	iface "red-packet/pkg/interface"
	util "red-packet/util/interface"
)

type service struct {
	repo      iface.IRepository
	cache     util.ICache
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
	Cache     util.ICache
	RedisLock *redislock.Client
}

func New(p Params) iface.IRedPacketService {
	return &service{
		repo:      p.Repo,
		cache:     p.Cache,
		redisLock: p.RedisLock,
	}
}
