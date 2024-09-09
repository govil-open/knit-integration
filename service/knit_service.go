package service

import (
	"encoding/json"
	"knit-integration/api/knit"
	"knit-integration/config"
	"knit-integration/httpclient"
	"knit-integration/logger"
	"knit-integration/model"

	"github.com/gin-gonic/gin"
)

type KnitService interface {
	GetPayroll(ctx *gin.Context) (model.Response, error)
}

type KnitServiceImpl struct {
	Conf *config.AppConfig
	Log  *logger.Logger
}

func NewKnitServiceImpl(c *config.AppConfig, logger *logger.Logger) *KnitServiceImpl {
	fs := &KnitServiceImpl{Conf: c, Log: logger}
	return fs
}

func (s *KnitServiceImpl) GetPayroll(ctx *gin.Context) (model.Response, error) {

	url := knit.EMPLOYEE_PAYROLL.GetAPI()
	headers := map[string]string{
		"accept":                "application/json",
		"Authorization":         s.Conf.KnitAPIKey,
		"X-Knit-Integration-Id": s.Conf.KnitIntegrationId,
	}
	params := map[string]string{
		"employeeId": ctx.Query("employeeId"),
		"month":      ctx.Query("month"),
	}

	body, err := httpclient.Get(url, headers, params)
	if err != nil {
		return model.Response{}, err
	}

	var response model.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return model.Response{}, err
	}

	return response, nil
}
