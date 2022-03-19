package service

import (
	"go.uber.org/fx"
	"red-packet/pkg/service/red_packet"
)

var Module = fx.Options(
	red_packet.Module,
)
