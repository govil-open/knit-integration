package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitializeDB(t *testing.T) {
	confojJSON, _ := LoadConfig("test")
	dbStore := InitializeDB(&confojJSON)
	require.NotNil(t, dbStore)
}
