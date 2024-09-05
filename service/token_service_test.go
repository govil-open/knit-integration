package service

import (
	"errors"
	"golang-boiler-plate/config"
	store "golang-boiler-plate/db"
	mockdb "golang-boiler-plate/db/mock"
	"golang-boiler-plate/logger"
	"golang-boiler-plate/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db := createMockStore(t)
	request := model.TokenRequest{
		Requestor:   "open",
		ExpiresIn:   300,
		Scope:       []string{},
		Authorities: []string{},
	}

	tokenService := new(TokenServiceImpl)
	db.EXPECT().CreateToken(&gin.Context{}, gomock.Any()).Return(uuid.New(), nil).Times(1)
	tokenResponse, _ := tokenService.Create(&gin.Context{}, db, &request)
	require.Equal(t, 300, tokenResponse.ExpiresIn)
}

func TestCreateError(t *testing.T) {
	mockStore := createMockStore(t)
	request := model.TokenRequest{
		Requestor:   "open",
		ExpiresIn:   300,
		Scope:       []string{},
		Authorities: []string{},
	}

	tokenService := new(TokenServiceImpl)
	mockStore.EXPECT().CreateToken(&gin.Context{}, gomock.Any()).Return(uuid.New(), errors.New("error")).Times(1)
	_, err := tokenService.Create(&gin.Context{}, mockStore, &request)
	require.Error(t, err)
}

func createMockStore(t *testing.T) *mockdb.MockStore {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	return mockdb.NewMockStore(mockCtrl)
}

func TestTokenServiceImpl_Create(t *testing.T) {
	request := model.TokenRequest{
		Requestor:   "open",
		ExpiresIn:   300,
		Scope:       []string{},
		Authorities: []string{},
	}
	mockDB := createMockStore(t)
	mockDB.EXPECT().CreateToken(&gin.Context{}, gomock.Any()).Return(uuid.New(), nil).Times(1)

	mockDBErr := createMockStore(t)
	mockDBErr.EXPECT().CreateToken(&gin.Context{}, gomock.Any()).Return(uuid.New(), errors.New("error")).Times(1)

	type fields struct {
		Store store.Store
		Conf  *config.AppConfig
		Log   *logger.Logger
	}
	type args struct {
		c  *gin.Context
		st store.Store
		tr *model.TokenRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.TokenResponse
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Store: mockDB,
				Conf:  configuration,
				Log:   logwrapper,
			},
			args: args{
				c:  &gin.Context{},
				st: mockDB,
				tr: &request,
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				Store: mockDB,
				Conf:  configuration,
				Log:   logwrapper,
			},
			args: args{
				c:  &gin.Context{},
				st: mockDBErr,
				tr: &request,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TokenServiceImpl{
				Store: tt.fields.Store,
				Conf:  tt.fields.Conf,
				Log:   tt.fields.Log,
			}
			_, err := s.Create(tt.args.c, tt.args.st, tt.args.tr)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenServiceImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
