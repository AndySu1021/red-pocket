package gin

import (
	"context"
	"errors"
	"go.uber.org/fx"
	"golang.org/x/time/rate"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Mode      string `mapstructure:"mode"`
	Port      string `mapstructure:"port"`
	RateLimit int    `mapstructure:"rate_limit"`
}

var limiter *rate.Limiter

func NewGin(lc fx.Lifecycle, cfg *Config) *gin.Engine {
	gin.SetMode(cfg.Mode)

	if cfg.RateLimit != 0 {
		limiter = initRateLimiter(cfg.RateLimit)
	}

	var e = gin.New()

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: e,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info().Msgf("Starting gin server, listen on %s.", cfg.Port)
			var err error
			go func() {
				if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Msgf("Fail to run gin server, err: %+v", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			defer log.Info().Msgf("Stopping gin HTTP server.")
			c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			return server.Shutdown(c)
		},
	})

	e.Use(LogRequest())
	e.Use(RequestIDMiddleware())
	e.Use(RateLimiterMiddleware())
	//e.Use(NewRecoverMiddleware())
	//e.Use(CORS())

	pprof.Register(e)
	return e
}

func initRateLimiter(rateLimit int) *rate.Limiter {
	return rate.NewLimiter(rate.Every(time.Duration(rateLimit)*time.Second), rateLimit * 2)
}

// RegisterDefaultRoute provide default handler
func RegisterDefaultRoute(e *gin.Engine) {
	e.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
}
