package restful

import (
	iface "red-packet/pkg/interface"
	ginTool "red-packet/util/gin"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RegisterAPIRouterParams struct {
	fx.In

	Engine   *gin.Engine
	Handlers []iface.IHandler `group:"handler"`
}

func RegisterAPIRouter(p RegisterAPIRouterParams) {
	p.Engine.Static("/public", "public")
	for _, h := range p.Handlers {
		h.Register(p.Engine.Group("/" + h.Version()))
	}

	ginTool.RegisterDefaultRoute(p.Engine)
}
