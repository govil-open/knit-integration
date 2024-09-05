package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTokenResponse(t *testing.T) {
	tokenResponse := NewTokenResponse("gsfgksgfbsfbkdbfks", 300, "bearer", "")
	require.Equal(t, tokenResponse.AccessToken, "gsfgksgfbsfbkdbfks")
}
