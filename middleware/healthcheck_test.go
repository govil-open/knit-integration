package middleware

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	handlerFunc := HealthCheck()
	require.NotNil(t, handlerFunc)
}

func TestNew(t *testing.T) {
	handlerFunc := New(Config{})
	require.NotNil(t, handlerFunc)
}
