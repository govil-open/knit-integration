package controller

import (
	"knit-integration/config"
	"knit-integration/middleware"
	"knit-integration/model"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
)

func getBaseRequest(c *gin.Context) (*model.BaseRequest, error) {
	authPayload, pok := c.Get(config.AuthPayloadKey)
	br := &model.BaseRequest{}
	if pok {
		m, ok := authPayload.(map[string]any)
		if !ok {
			log.Error("authpayload not present")
			return br, &middleware.ValidationError{Message: "Auth payload not present"}
		}
		companyID, k := m[config.CompanyID].(float64)
		if !k {
			log.Error("companyID not present")
			return br, &middleware.ValidationError{Message: "CompanyID not present"}
		}
		accountID, k := m[config.AccountID].(float64)
		if !k {
			log.Error("accountID not present")
			return br, &middleware.ValidationError{Message: "AccountID not present"}
		}
		userID, k := m[config.UserID].(float64)
		if !k {
			log.Error("userID not present")
			return br, &middleware.ValidationError{Message: "UserID not present"}
		}
		br.UserID = int64(userID)
		br.CompanyID = int64(companyID)
		br.AccountID = int64(accountID)
		return br, nil
	}
	return br, &middleware.ValidationError{Message: "Auth payload not present"}
}

func getApplicationError(httpStatusCode int, errorCode int, message string, err error, l *i18n.Localizer) middleware.ApplicationError {
	tm := config.Translate(&message, l)
	if err == nil {
		return middleware.HandleErrors(httpStatusCode, errorCode, tm)
	}
	return middleware.HandleValidationErrors(err, httpStatusCode, errorCode, tm, l)
}
