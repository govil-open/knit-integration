package service

import (
	"golang-boiler-plate/config"
	store "golang-boiler-plate/db"
	db "golang-boiler-plate/db/sqlc"
	"golang-boiler-plate/logger"
	"golang-boiler-plate/model"
	"golang-boiler-plate/util"

	"github.com/gin-gonic/gin"
)

type TokenService interface {
	Create(ctx *gin.Context, store store.Store, tokenRequest *model.TokenRequest) (*model.TokenResponse, error)
}

type TokenServiceImpl struct {
	Store store.Store
	Conf  *config.AppConfig
	Log   *logger.Logger
}

func NewTokenServiceImpl(s store.Store, c *config.AppConfig, logger *logger.Logger) *TokenServiceImpl {
	fs := &TokenServiceImpl{Store: s, Conf: c, Log: logger}
	return fs
}

func (s *TokenServiceImpl) Create(c *gin.Context, st store.Store, tr *model.TokenRequest) (*model.TokenResponse, error) {
	tokenString, err := util.RandomString(config.TokenLength)
	if err != nil {
		s.Log.Error(&tr.BaseRequest, err)
		return nil, err
	}
	arg := db.CreateTokenParams{
		Audience:    tr.Requestor,
		Token:       tokenString,
		ExpiresIn:   int32(tr.ExpiresIn),
		Scope:       tr.Scope,
		Authorities: tr.Authorities,
	}
	token, err := st.CreateToken(c, arg)
	if err != nil {
		s.Log.Error(&tr.BaseRequest, err)
		return nil, err
	}
	return model.NewTokenResponse(tokenString, tr.ExpiresIn, config.TokenType, token.String()), nil
}
