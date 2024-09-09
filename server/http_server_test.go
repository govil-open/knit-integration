package server

import (
	"knit-integration/config"
	"knit-integration/db/mock"
	"knit-integration/logger"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/require"
)

var (
	logwrapper *logger.Logger
)

func TestInit(t *testing.T) {
	configJSON := config.AppConfig{
		GinMode:  "test",
		DBDriver: "postgres",
	}
	ctrl := gomock.NewController(t)
	relicApp, _ := newrelic.NewApplication(
		newrelic.ConfigAppName(""),
		newrelic.ConfigLicense(""),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	_, err := Init(&configJSON, mock.NewMockStore(ctrl), relicApp, logwrapper)
	require.NoError(t, err)
}
