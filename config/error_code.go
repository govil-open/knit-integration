package config

import "errors"

const (
	BadRequestErrorCode = 411
	ServerErrorCode     = 511
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)
