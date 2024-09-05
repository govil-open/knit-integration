package service

import (
	"encoding/json"
	"golang-boiler-plate/api/knit"
	"golang-boiler-plate/config"
	"golang-boiler-plate/logger"
	"golang-boiler-plate/model"
	"io"
	"net/http"

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
	employeeId := ctx.Query("employeeId")
	month := ctx.Query("month") //format - yyyy-mm

	payrollAPI := knit.EMPLOYEE_PAYROLL.GetAPI()
	url := payrollAPI + "?employeeId=" + employeeId + "&month=" + month

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Response{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", s.Conf.KnitAPIKey)
	req.Header.Add("X-Knit-Integration-Id", s.Conf.KnitIntegrationId)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.Response{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.Response{}, err
	}

	var response model.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return model.Response{}, err
	}

	return response, nil
}
