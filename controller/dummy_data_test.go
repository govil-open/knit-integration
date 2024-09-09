package controller

import (
	"encoding/json"
	"knit-integration/model"
)

func createDummyAuthMap() map[string]any {
	m := make(map[string]any)
	m["jti"] = "yAmWssrLCVUR6fev"
	m["mobile_number"] = "5435353435353"
	m["users_id"] = float64(357263)
	m["accounts_id"] = float64(353793)
	m["companies_id"] = float64(353795)
	m["sub"] = 357263
	m["iss"] = "http://127.0.0.1:8000/api/users/login"
	m["iat"] = 1663572509
	m["exp"] = 1723572509
	m["nbf"] = 1663572509
	return m
}

func createTokenStruct() model.TokenRequest {
	req := model.TokenRequest{
		ExpiresIn:   8,
		Requestor:   "jbk",
		Scope:       []string{"jbcu", "jbk"},
		Authorities: []string{"wcihw", "ncoj"},
	}
	return req
}

func createTokenBody() []byte {
	req := createTokenStruct()
	body, _ := json.Marshal(req)
	return body
}

func createBindErrBody() []byte {
	req := "wjbbhuc"
	body, _ := json.Marshal(req)
	return body
}
