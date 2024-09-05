package model

type TokenRequest struct {
	ExpiresIn   int      `json:"expires_in"  binding:"required"`
	Requestor   string   `json:"requestor"  binding:"required"`
	Scope       []string `json:"scope"`
	Authorities []string `json:"authorities"`
	BaseRequest
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Jti         string `json:"jti"`
}

func NewTokenResponse(accessToken string, expirestIn int, tokenType string, jti string) *TokenResponse {
	tokenResponse := &TokenResponse{
		AccessToken: accessToken,
		TokenType:   tokenType,
		ExpiresIn:   expirestIn,
		Jti:         jti,
	}
	return tokenResponse
}
