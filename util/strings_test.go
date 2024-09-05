package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertNewline(t *testing.T) {
	output := ConvertNewline("token\r\ntoken\rtoken\n", "\n")
	require.Equal(t, output, "token\ntoken\ntoken\n")
	output = ConvertNewline("token\r\ntoken\rtoken\n", "")
	require.Equal(t, output, "tokentokentoken")
}

func TestSuffix(t *testing.T) {
	output := Suffix("user", ".")
	require.Equal(t, output, "user.")
	output = Suffix("", ".")
	require.Equal(t, output, "")
}
