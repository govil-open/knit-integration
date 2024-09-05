package knit

type KnitAPI int

type APIType int

type API struct {
	URL  string
	Type APIType
}

const (
	GET APIType = iota
	POST
	PUT
	DELETE
)

const (
	EMPLOYEE_PAYROLL KnitAPI = iota
	LEAVE_REQUEST
	LEAVE_BALANCE
	CREATE_LEAVE_REQUEST
	ATTENANCE
)

var apis = [...]API{
	{URL: "https://api.getknit.dev/v1.0/hr.employees.payroll.get", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.leave.requests", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.leave.balance", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.leave.create", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.attendance", Type: GET},
}

func (knitAPI KnitAPI) GetAPI() string {
	return apis[knitAPI].URL
}

func (knitAPI KnitAPI) GetType() APIType {
	return apis[knitAPI].Type
}

func (a APIType) String() string {
	return [...]string{"GET", "POST", "PUT", "DELETE"}[a]
}
