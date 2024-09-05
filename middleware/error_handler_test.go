package middleware

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

var (
	validationFailed = "validation failed"
)

func TestHandleValidationErrors(t *testing.T) {
	handleError := HandleValidationErrors(errors.New(validationFailed), 400, 411, validationFailed, conf.Localizer)
	require.Nil(t, handleError.FieldErrors)
	handleError = HandleValidationErrors(nil, 400, 411, validationFailed, conf.Localizer)
	require.Nil(t, handleError.FieldErrors)
}

type TokenRequest struct {
	AccessRequest string `validate:"required"`
	Requester     string `validate:"required"`
}

func TestHandleValidationErrorFailedValidation(t *testing.T) {
	tokenRequest := TokenRequest{}
	validator := validator.New()
	err := validator.Struct(tokenRequest)
	handleError := HandleValidationErrors(err, 400, 411, validationFailed, conf.Localizer)
	require.NotNil(t, handleError.FieldErrors)
}

func TestHandleErrors(t *testing.T) {
	handleError := HandleErrors(400, 411, validationFailed)
	require.Nil(t, handleError.FieldErrors)
	require.Equal(t, handleError.Message, validationFailed)
}

func TestMessageForTagRequired(t *testing.T) {
	message := messageForTag("required", conf.Localizer)
	require.Equal(t, message, "is required")
	message = messageForTag("email", conf.Localizer)
	require.Equal(t, message, "")
}

func TestMessageForTagUnsupported(t *testing.T) {
	message := messageForTag("access_request_types", conf.Localizer)
	require.Equal(t, message, "is unsupported")
}
