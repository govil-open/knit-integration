package model

import "time"

type LeaveTypeEnum string
type IsEligible string

type LeaveRequest struct {
	ID          string    `json:"id"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	RequestedOn time.Time `json:"requestedOn"`
	Note        string    `json:"note"`
	Status      Status    `json:"status"`
	LeaveType   LeaveType `json:"leaveType"`
	Unit        Unit      `json:"unit"`
	Amount      float64   `json:"amount"`
	IsPaid      IsPaid    `json:"isPaid"`
}

const (
	VACATION            LeaveTypeEnum = "VACATION"
	SICK                LeaveTypeEnum = "SICK"
	PERSONAL            LeaveTypeEnum = "PERSONAL"
	JURY_DUTY           LeaveTypeEnum = "JURY_DUTY"
	VOLUNTEER           LeaveTypeEnum = "VOLUNTEER"
	BEREAVEMENT         LeaveTypeEnum = "BEREAVEMENT"
	LEAVE_NOT_SPECIFIED LeaveTypeEnum = "NOT_SPECIFIED"
)

// This struct can be constructed as an enum with multiple values
// Once we have clarity we can re-struct this one.
type LeaveType struct {
	ID   string
	Name string
	Type LeaveTypeEnum
}

type LeaveBalance struct {
	LeaveType  LeaveType  `json:"leaveType"`
	Unit       Unit       `json:"unit"`
	Balance    float64    `json:"balance"`
	LeaveUsed  float64    `json:"used"`
	IsPaid     IsPaid     `json:"isPaid"`
	IsEligible IsEligible `json:"isEligible"`
}

const (
	IS_ELIGIBLE              IsEligible = "TRUE"
	IS_NOT_ELIGIBLE          IsEligible = "FALSE"
	LEAVE_TYPE_NOT_SPEFICIED IsEligible = "NOT_SPECIFIED"
)
