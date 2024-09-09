package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"knit-integration/config"
	"knit-integration/logger"
	"knit-integration/util"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const (
	AuthFieldLength = 2
)

func AuthMiddleware(conf *config.AppConfig, log *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(config.AuthHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, getError("auth_header_not_found", nil, conf.Localizer))
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < AuthFieldLength {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, getError("auth_header_format", nil, conf.Localizer))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != config.AuthHeaderType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, getError("auth_header_unsupported", nil, conf.Localizer))
			return
		}
		accessToken := fields[1]
		payload, err := util.Post(ctx, conf, log, "application/json", getAuthRequestBody(accessToken))
		fmt.Println(err)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, getError("auth_token_invalid", err, conf.Localizer))
			return
		}
		ctx.Set(config.AuthPayloadKey, payload)
		ctx.Next()
	}
}

func getError(key string, err error, l *i18n.Localizer) ApplicationError {
	tm := config.Translate(&key, l)
	return HandleValidationErrors(err, http.StatusUnauthorized, config.UnauthorizedErrorCode, tm, l)
}

func getAuthRequestBody(accessToken string) map[string]any {
	body := make(map[string]any)
	body["token"] = accessToken
	return body
}
