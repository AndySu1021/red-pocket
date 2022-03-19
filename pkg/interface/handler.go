package iface

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type IHandler interface {
	Version() string
	Register(g *gin.RouterGroup)
}

type HandlerOut struct {
	fx.Out

	Handler IHandler `group:"handler"`
}
