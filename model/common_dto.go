package model

type BaseRequest struct {
	CompanyID int64 `json:"companies_id"`
	AccountID int64 `json:"accounts_id"`
	UserID    int64 `json:"users_id"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
