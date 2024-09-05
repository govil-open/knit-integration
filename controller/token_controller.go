package controller

import (
	"golang-boiler-plate/config"
	"golang-boiler-plate/model"
	"golang-boiler-plate/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	service *service.TokenServiceImpl
	conf    *config.AppConfig
}

func NewTokenController(s *service.TokenServiceImpl, config *config.AppConfig) *TokenController {
	return &TokenController{service: s, conf: config}
}

// CreateToken godoc
// @Summary      Create Token
// @Description  Create Token for authentication/authorization
// @Accept       json
// @Produce      json
// @Param        model.TokenRequest  body    model.TokenRequest  true  "Token Request"
// @Success      201                 object  model.TokenResponse
// @Failure      400                 object  middleware.ApplicationError
// @Failure      500                 object  middleware.ApplicationError
// @Router       /v1/tokens [post].
func (u *TokenController) Create(c *gin.Context) {
	input := model.TokenRequest{}
	userDet, _ := getBaseRequest(c)
	if err := c.ShouldBindJSON(&input); err != nil {
		e := getApplicationError(http.StatusBadRequest, config.BadRequestErrorCode, "validation_failed", err, u.conf.Localizer)
		c.JSON(http.StatusBadRequest, e)
		return
	}
	input.BaseRequest = *userDet
	tokenResponse, err := u.service.Create(c, u.service.Store, &input)
	if err != nil {
		e := getApplicationError(http.StatusInternalServerError, config.ServerErrorCode, "server_error", nil, u.conf.Localizer)
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	c.JSON(http.StatusCreated, tokenResponse)
}
