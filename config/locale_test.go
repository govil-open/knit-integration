package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranslate(t *testing.T) {
	message := "validation_failed"
	checkForMessageNotEmpty(t, &message)
	message = "required"
	checkForMessageNotEmpty(t, &message)
	message = "abcdef"
	translatedMessage := Translate(&message, config.Localizer)
	require.Equal(t, translatedMessage, "")
}

func checkForMessageNotEmpty(t *testing.T, message *string) {
	translatedMessage := Translate(message, config.Localizer)
	require.NotEmpty(t, translatedMessage)
}
