package config

import (
	"golang-boiler-plate/logger"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetupSentry(t *testing.T) {
	configJSON, _ := LoadConfig("test")
	logwrapper := logger.Init()
	err := setUpSentry(&configJSON, logwrapper)
	require.NoError(t, err)
}
