package enum

type Status string
type Unit string
type IsPaid string
type Encoding string

const (
	REQUESTED     Status = "REQUESTED"
	APPROVED      Status = "APPROVED"
	DECLINED      Status = "DECLINED"
	CANCELLED     Status = "CANCELLED"
	DELETED       Status = "DELETED"
	NOT_SPECIFIED Status = "NOT_SPECIFIED"
	ACTIVE        Status = "ACTIVE"
	INACTIVE      Status = "INACTIVE"

	DAYS  Unit = "DAYS"
	HOURS Unit = "HOURS"

	TRUE  IsPaid = "TRUE"
	FALSE IsPaid = "FALSE"

	//Attendance related status
	PRESENT    Status = "PRESENT"
	ABSENT     Status = "ABSENT"
	HOLIDAY    Status = "HOLIDAY"
	WEEKLY_OFF Status = "WEEKLY_OFF"

	Base64                 Encoding = "BASE64"
	ENCODING_NOT_SPECIFIED Encoding = "NOT_SPECIFIED"
)
