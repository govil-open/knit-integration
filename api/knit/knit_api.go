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
	LIST_ALL_LEAVE_TYPES
	LEAVE_REQUEST
	LEAVE_BALANCE
	CREATE_LEAVE_REQUEST
	ATTENDANCE
	GET_COMPENSATION_PLANS
	UPDATE_EMPLOYEE_COMPENSATION
	EMPLOYEE_DOCUMENT
	UPDATE_EMPLOYEE_DOCUMENT
	DOCUMENT_CATEGORIES
	UPDATE_EMPLOYEE
	CREATE_EMPLOYEE
	GET_POSITION_DETAILS
	LIST_OF_VALUES
	TERMINATE_EMPLOYEE
	TERMINATION_REASON
	GET_COMPANY_LIST
	LIST_COMPANY_WIDE_DEDUCTION
	CREATE_COMPANY_WIDE_DEDUCTION
	UPDATE_COMPANY_WIDE_DEDUCTION
	ENROLL_EMPLOYEE_IN_DEDUDCTION
	UNENROLL_EMPLOYEE_IN_DEDUDCTION
	UPDATE_EMPLOYEE_DEDUCTION
)

var apis = [...]API{
	//Payroll API
	{URL: "https://api.getknit.dev/v1.0/hr.employees.payroll.get", Type: GET},

	//Leave APIs
	{URL: "https://api.getknit.dev/v1.0/hr.leave.types", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.leave.requests", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.leave.balance", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.leave.create", Type: POST},

	//Attendance API
	{URL: "https://api.getknit.dev/v1.0/hr.employees.attendance", Type: GET},

	//Compensation APIs
	{URL: "https://api.getknit.dev/v1.0/hr.compensation.plans", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employee.compensation.update", Type: POST},

	//Document APIs
	{URL: "https://api.getknit.dev/v1.0/hr.employees.documents", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.employees.document.upload", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.documents.categories", Type: GET},

	//Employee APIs
	{URL: "https://api.getknit.dev/v1.0/hr.employee.update", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.employee.create", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.documents.categories", Type: GET},

	//List of values API
	{URL: "https://api.getknit.dev/v1.0/hr.employees.field.values", Type: GET},

	//Termination APIs
	{URL: "https://api.getknit.dev/v1.0/hr.employee.terminate", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.termination.reasons", Type: GET},

	//Company APIs
	{URL: "https://api.getknit.dev/v1.0/hr.companies.list", Type: GET},

	//Deduction APIs
	{URL: "https://api.getknit.dev/v1.0/hr.deductions.list", Type: GET},
	{URL: "https://api.getknit.dev/v1.0/hr.deductions.create", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.deductions.update", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.employee.deductions.enroll", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.employee.deductions.unenroll", Type: POST},
	{URL: "https://api.getknit.dev/v1.0/hr.employee.deductions.update", Type: POST},
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
