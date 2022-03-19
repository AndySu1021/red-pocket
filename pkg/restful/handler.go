package restful

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	iface "red-packet/pkg/interface"
)

type handler struct {
	redPacketSvc iface.IRedPacketService
}

func (h *handler) Version() string {
	return "/api/v1"
}

func (h *handler) Register(g *gin.RouterGroup) {
	g.POST("/red-packet/send", h.Send)
	g.POST("/red-packet/grab", h.Grab)
}

var Module = fx.Options(
	fx.Provide(
		NewHandler,
	),
)

type Params struct {
	fx.In

	RedPacketSvc iface.IRedPacketService
}

func NewHandler(p Params) iface.HandlerOut {
	return iface.HandlerOut{
		Handler: &handler{
			redPacketSvc: p.RedPacketSvc,
		},
	}
}
