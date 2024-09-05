package model

import "time"

type SalaryType string

const (
	// Earnings
	INCOME            SalaryType = "INCOME"
	OVERTIME          SalaryType = "OVERTIME"
	HRA               SalaryType = "HRA"
	BASIC             SalaryType = "BASIC"
	SPECIAL_ALLOWANCE SalaryType = "SPECIAL_ALLOWANCE"
	LTA               SalaryType = "LTA"

	// Contributions
	PF                SalaryType = "PF"
	MEDICAL_INSURANCE SalaryType = "MEDICAL_INSURANCE"
	_401K             SalaryType = "401k"
	LIFE_INSURANCE    SalaryType = "LIFE_INSURANCE"
	VISION            SalaryType = "VISION"
	DENTAL            SalaryType = "DENTAL"

	// Deductions
	PROF_TAX    SalaryType = "PROF_TAX"
	TDS         SalaryType = "TDS"
	FEDERAL_TAX SalaryType = "FEDERAL_TAX"
	STATE_TAX   SalaryType = "STATE_TAX"
	TAX         SalaryType = "TAX"
	HSA         SalaryType = "HSA"
)

type SalaryItem struct {
	SalaryType SalaryType `json:"type"`
	Amount     float64    `json:"amount"`
}

type Payroll struct {
	EmployeeID         string       `json:"employeeId"`
	GrossPay           float64      `json:"grossPay"`
	NetPay             float64      `json:"netPay"`
	Earnings           []SalaryItem `json:"earnings"`
	Contributions      []SalaryItem `json:"contributions"`
	Deductions         []SalaryItem `json:"deductions"`
	ProcessedDate      time.Time    `json:"processedDate"`
	PayDate            time.Time    `json:"payDate"`
	PayPeriodStartDate time.Time    `json:"payPeriodStartDate"`
	PayPeriodEndDate   time.Time    `json:"payPeriodEndDate"`
}
