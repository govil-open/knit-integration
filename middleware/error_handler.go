package middleware

import (
	"errors"
	"fmt"
	"knit-integration/config"

	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ApplicationError struct {
	Message     string       `json:"message"`
	ErrorCode   int          `json:"error_code"`
	StatusCode  int          `json:"status_code"`
	FieldErrors []FieldError `json:"errors"`
}

type FieldError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func HandleValidationErrors(err error, statusCode int, errorCode int, message string, l *i18n.Localizer) ApplicationError {
	var validationError validator.ValidationErrors
	var fieldErrors []FieldError
	if errors.As(err, &validationError) {
		fieldErrors = make([]FieldError, len(validationError))
		for i, fe := range validationError {
			fieldErrors[i] = FieldError{fe.Field(), messageForTag(fe.Tag(), l)}
		}
	}
	return ApplicationError{StatusCode: statusCode, ErrorCode: errorCode, Message: message, FieldErrors: fieldErrors}
}

func HandleErrors(statusCode int, errorCode int, message string) ApplicationError {
	return ApplicationError{StatusCode: statusCode, ErrorCode: errorCode, Message: message, FieldErrors: nil}
}

func messageForTag(tag string, l *i18n.Localizer) string {
	var message string
	switch tag {
	case "required":
		requiredMessage := "required"
		message = config.Translate(&requiredMessage, l)
	case "access_request_types":
		unsupportedMessage := "unsupported"
		message = config.Translate(&unsupportedMessage, l)
	}
	return message
}

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%v", v.Message)
}
