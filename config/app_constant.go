package config

import "net/http"

// put all the application constants here.
const (
	AuthHeaderKey         = "Authorization"
	AuthHeaderType        = "bearer"
	AuthPayloadKey        = "auth_payload"
	TokenIssuer           = "token-service"
	TokenType             = "bearer"
	TokenLength           = 30
	DefaultHeaderName     = "X-Health-Check"
	DefaultHeaderValue    = "1"
	DefaultResponseCode   = http.StatusOK
	CompanyID             = "companies_id"
	AccountID             = "accounts_id"
	UserID                = "users_id"
	HTTPSuccessCode       = 200
	UnauthorizedErrorCode = 401
)
