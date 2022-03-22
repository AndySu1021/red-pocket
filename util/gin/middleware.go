package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"strings"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !strings.Contains(c.Request.URL.Path, "health") {
			log.Ctx(c.Request.Context()).Info().
				Str("url", c.Request.URL.String()).
				Str("method", c.Request.Method).
				Interface("header", c.Request.Header).
				Bool("access_log", true).
				Msgf("access log")
		}
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = xid.New().String()
		}

		logger := log.With().Str("trace_id", traceID).Logger()

		deviceID := c.GetHeader("X-Device-ID")

		ctx := context.WithValue(c.Request.Context(), "X-Trace-ID", traceID)
		ctx = context.WithValue(ctx, "X-Device-ID", deviceID)
		ctx = logger.WithContext(ctx)

		c.Request = c.Request.WithContext(ctx)
		c.Request.Header.Set("X-Trace-ID", traceID)
		c.Writer.Header().Set("X-Trace-ID", traceID)
		c.Next()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter == nil {
			c.Next()
		}

		err := limiter.Wait(c.Request.Context())
		if err != nil {
			return
		}

		c.Next()
	}
}
