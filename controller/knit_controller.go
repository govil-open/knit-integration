package controller

import (
	"knit-integration/config"
	"knit-integration/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KnitController struct {
	Config  *config.AppConfig
	Service *service.KnitServiceImpl
}

func NewKnitController(s *service.KnitServiceImpl, config *config.AppConfig) *KnitController {
	return &KnitController{Service: s, Config: config}
}

func (u *KnitController) GetPayroll(c *gin.Context) {
	res, err := u.Service.GetPayroll(c)
	if err != nil {
		e := getApplicationError(http.StatusInternalServerError, config.ServerErrorCode, "server_error", err, u.Config.Localizer)
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	c.JSON(http.StatusOK, res)
}
