package controller

import (
	"bytes"
	"knit-integration/config"
	mockdb "knit-integration/db/mock"
	"knit-integration/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func createMockStore(t *testing.T) *mockdb.MockStore {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	return mockdb.NewMockStore(mockCtrl)
}

func createContext() *gin.Context {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	authPayload := createDummyAuthMap()
	c.Set(config.AuthPayloadKey, authPayload)
	return c
}

func TestTokenController_Create(t *testing.T) {
	c := createContext()
	body := createTokenBody()
	c.Request = httptest.NewRequest(http.MethodPost, "/tokens", bytes.NewBuffer(body))
	c.Request.Header.Add("Content-Type", "application/json")
	dbStore := createMockStore(t)
	id := uuid.New()
	dbStore.EXPECT().CreateToken(c, gomock.Any()).Return(id, nil).Times(1)
	ctxBindErr := createContext()

	reqBody := createBindErrBody()
	ctxBindErr.Request = httptest.NewRequest(http.MethodPost, "/tokens", bytes.NewBuffer(reqBody))
	ctxBindErr.Request.Header.Add("Content-Type", "application/json")

	tokenService := &service.TokenServiceImpl{
		Store: dbStore,
		Conf:  configuration,
		Log:   logwrapper,
	}

	type fields struct {
		service *service.TokenServiceImpl
		conf    *config.AppConfig
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Create - Success",
			fields: fields{
				service: tokenService,
				conf:    configuration,
			},
			args: args{
				c: c,
			},
		},
		{
			name: "Create - BindError",
			fields: fields{
				conf: configuration,
			},
			args: args{
				c: ctxBindErr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &TokenController{
				service: tt.fields.service,
				conf:    tt.fields.conf,
			}
			u.Create(tt.args.c)
		})
	}
}
