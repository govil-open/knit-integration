package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	randomStringsMap := make(map[string]bool)
	for i := 0; i < 100; i++ {
		data, _ := RandomString(10)
		if randomStringsMap[data] == true {
			require.Fail(t, "strings are not random")
		}
		randomStringsMap[data] = true
	}
}

func TestRandomStringNotEmpty(t *testing.T) {
	randomString, _ := RandomString(10)
	require.NotEmpty(t, randomString)
}

func TestUuid(t *testing.T) {
	uuid := UUID()
	require.NotEmpty(t, uuid)
}
