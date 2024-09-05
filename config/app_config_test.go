package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	configJSON, err := LoadConfig("test")
	require.NoError(t, err)
	require.NotEmpty(t, configJSON.HTTPServerAddress)
	require.NotEmpty(t, configJSON.GinMode)
}
