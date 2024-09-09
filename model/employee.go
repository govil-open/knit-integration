package model

import (
	"knit-integration/model/enum"
	"time"
)

// model id is employee_profile
type EmployeeProfile struct {
	FirstName        string              `json:"firstName"`
	LastName         string              `json:"lastName"`
	ID               string              `json:"id"`
	EmployeeNumber   string              `json:"employeeNumber"`
	WorkEmail        string              `json:"workEmail"`
	StartDate        time.Time           `json:"startDate"`
	BirthDate        time.Time           `json:"birthDate"`
	TerminationDate  *time.Time          `json:"terminationDate"`
	EmploymentStatus enum.Status         `json:"employmentStatus"`
	MaritalStatus    enum.MaritalStatus  `json:"maritalStatus"`
	Gender           enum.Gender         `json:"gender"`
	EmploymentType   enum.EmploymentType `json:"employmentType"`
}

type Phone struct {
	Number string         `json:"number"`
	Type   enum.PhoneType `json:"type"`
}

// model id is employee_contactInfo
type Contact struct {
	PersonalEmails []string `json:"personalEmails"`
	Phones         []Phone  `json:"phones"`
}

// model id is employee_orgStructure
type EmployeeOrgStructure struct {
	Designation string          `json:"designation"`
	Department  string          `json:"department"`
	Manager     EmployeeManager `json:"manager"`
}

type EmployeeManager struct {
	ID        string `json:"id"`
	WorkEmail string `json:"workEmail"`
}

// model id is employee_dependent
type EmployeeDependents struct {
	Dependents []EmployeeDependent `json:"dependents"`
}

type EmployeeDependent struct {
	ID        string        `json:"id"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Relation  enum.Relation `json:"relation"`
	BirthDate time.Time     `json:"birthDate"`
	Gender    enum.Gender   `json:"gender"`
}

// model id is employee_location
type EmployeeLocation struct {
	WorkAddress      Location `json:"workAddress"`
	PermanentAddress Location `json:"permanentAddress"`
	PresentAddress   Location `json:"presentAddress"`
}

type Location struct {
	AddressLine1 string           `json:"addressLine1"`
	AddressLine2 string           `json:"addressLine2"`
	City         string           `json:"city"`
	State        string           `json:"state"`
	Country      enum.Country     `json:"country"`
	ZipCode      string           `json:"zipCode"`
	AddressType  enum.AddressType `json:"addressType"`
}

// model id is employee_bankAccount
type RoutingInfo struct {
	Type   enum.RoutingType `json:"type"`
	Number string           `json:"number"`
}

type BankAccount struct {
	AccountNumber string           `json:"accountNumber"`
	AccountType   enum.AccountType `json:"accountType"`
	BankName      string           `json:"bankName"`
	RoutingInfo   []RoutingInfo    `json:"routingInfo"`
	PayTypes      []enum.PayType   `json:"payTypes"`
}

type EmployeeBankAccount struct {
	BankAccounts []BankAccount `json:"bankAccounts"`
}

type IDData struct {
	Type                 enum.IDType `json:"type"`
	SubType              string      `json:"subType"`
	IdentificationNumber string      `json:"identificationNumber"`
}

// model id is employee_identification
type EmployeeIdentification struct {
	EmployeeIdentificationData []IDData `json:"employeeIdentificationData"`
}

// model id is employee_kycDocuments
type EmployeeKYC struct {
	AadharNumber string `json:"aadharNumber"`
	AadharURL    string `json:"aadharURL"`
	PanNumber    string `json:"panNumber"`
	PanURL       string `json:"panURL"`
}
