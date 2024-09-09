package controller

import (
	"knit-integration/config"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetApplicationError(t *testing.T) {
	ae := getApplicationError(200, 211, "", nil, configuration.Localizer)
	require.Equal(t, ae.StatusCode, 200)
}

func TestGetBaseRequest(t *testing.T) {
	http := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(http)
	authPayload := createDummyAuthMap()

	c.Set(config.AuthPayloadKey, authPayload)
	br, err := getBaseRequest(c)
	require.Empty(t, err)
	require.Equal(t, br.CompanyID, int64(353795))
	require.Equal(t, br.AccountID, int64(353793))
	require.Equal(t, br.UserID, int64(357263))
}
