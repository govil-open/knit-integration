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

type Status string
type Unit string
type IsPaid string

const (
	REQUESTED     Status = "REQUESTED"
	APPROVED      Status = "APPROVED"
	DECLINED      Status = "DECLINED"
	CANCELLED     Status = "CANCELLED"
	DELETED       Status = "DELETED"
	NOT_SPECIFIED Status = "NOT_SPECIFIED"

	DAYS               Unit = "DAYS"
	HOURS              Unit = "HOURS"
	UNIT_NOT_SPECIFIED Unit = "NOT_SPECIFIED"

	TRUE               IsPaid = "TRUE"
	FALSE              IsPaid = "FALSE"
	PAID_NOT_SPECIFIED IsPaid = "NOT_SPECIFIED"
)
