package model

import (
	"knit-integration/model/enum"
	"time"
)

type LeaveTypeEnum string
type IsEligible string

type LeaveRequest struct {
	ID          string      `json:"id"`
	StartDate   time.Time   `json:"startDate"`
	EndDate     time.Time   `json:"endDate"`
	RequestedOn time.Time   `json:"requestedOn"`
	Note        string      `json:"note"`
	Status      enum.Status `json:"status"`
	LeaveType   LeaveType   `json:"leaveType"`
	Unit        enum.Unit   `json:"unit"`
	Amount      float64     `json:"amount"`
	IsPaid      enum.IsPaid `json:"isPaid"`
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

type LeaveType struct {
	ID   string        `json:"id"`
	Name string        `json:"name"`
	Type LeaveTypeEnum `json:"type"`
}

type LeaveBalance struct {
	LeaveType  LeaveType   `json:"leaveType"`
	Unit       enum.Unit   `json:"unit"`
	Balance    float64     `json:"balance"`
	LeaveUsed  float64     `json:"used"`
	IsPaid     enum.IsPaid `json:"isPaid"`
	IsEligible IsEligible  `json:"isEligible"`
}

const (
	IS_ELIGIBLE              IsEligible = "TRUE"
	IS_NOT_ELIGIBLE          IsEligible = "FALSE"
	LEAVE_TYPE_NOT_SPEFICIED IsEligible = "NOT_SPECIFIED"
)
