package iface

import (
	"context"
	"red-packet/pkg/model/dto"
	"red-packet/pkg/model/option"
)

type IRedPacketService interface {
	Send(ctx context.Context, activity *dto.Activity) (bool, error)
	Grab(ctx context.Context, redPacketOpt *option.RedPacketOption, userOpt *option.UserOption) (uint64, error)
}
