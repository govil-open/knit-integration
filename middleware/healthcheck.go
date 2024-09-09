package middleware

import (
	"knit-integration/config"

	"github.com/gin-gonic/gin"
)

// Config holds the configuration values.
type Config struct {
	HeaderName   string
	HeaderValue  string
	ResponseCode int
}

// Default creates a new middileware with the default configuration.
func HealthCheck() gin.HandlerFunc {
	return New(Config{
		HeaderName:   config.DefaultHeaderName,
		HeaderValue:  config.DefaultHeaderValue,
		ResponseCode: config.DefaultResponseCode,
	})
}

// New creates a new middileware with the `cfg`.
func New(cfg Config) gin.HandlerFunc {
	if cfg.HeaderName == "" {
		cfg.HeaderName = config.DefaultHeaderName
	}
	if cfg.HeaderValue == "" {
		cfg.HeaderValue = config.DefaultHeaderValue
	}
	if cfg.ResponseCode == 0 {
		cfg.ResponseCode = config.DefaultResponseCode
	}

	return func(ctx *gin.Context) {
		if ctx.GetHeader(cfg.HeaderName) == cfg.HeaderValue {
			ctx.String(cfg.ResponseCode, "")
			ctx.Abort()
		}
		ctx.Next()
	}
}
