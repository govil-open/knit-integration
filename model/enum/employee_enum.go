package enum

type MaritalStatus string
type Gender string
type Relation string
type EmploymentType string
type Country string
type AddressType string
type PhoneType string
type RoutingType string
type AccountType string
type PayType string
type IDType string

const (
	MARRIED                      MaritalStatus = "MARRIED"
	SINGLE                       MaritalStatus = "SINGLE"
	WIDOWED                      MaritalStatus = "WIDOWED"
	SEPARATED                    MaritalStatus = "SEPARATED"
	DIVORCED                     MaritalStatus = "DIVORCED"
	MARITAL_STATUS_NOT_SPECIFIED MaritalStatus = "NOT_SPECIFIED"

	RELATION_NOT_SPECIFIED Relation = "NOT_SPECIFIED"
	PARENT                 Relation = "PARENT"
	FATHER                 Relation = "FATHER"
	MOTHER                 Relation = "MOTHER"
	SON                    Relation = "SON"
	DAUGHTER               Relation = "DAUGHTER"
	CHILD                  Relation = "CHILD"
	BROTHER                Relation = "BROTHER"
	SISTER                 Relation = "SISTER"
	SIBLING                Relation = "SIBLING"
	HUSBAND                Relation = "HUSBAND"
	WIFE                   Relation = "WIFE"
	SPOUSE                 Relation = "SPOUSE"

	MALE                 Gender = "MALE"
	FEMALE               Gender = "FEMALE"
	GENDER_NOT_SPECIFIED Gender = "NOT_SPECIFIED"
	NON_BINARY           Gender = "NON_BINARY"

	FULL_TIME                       EmploymentType = "FULL_TIME"
	PART_TIME                       EmploymentType = "PART_TIME"
	CONTRACT                        EmploymentType = "CONTRACT"
	EMPLOYEEMENT_TYPE_NOT_SPECIFIED EmploymentType = "NOT_SPECIFIED"

	WORK_ADDRESS   AddressType = "WORK"
	HOME_PRESENT   AddressType = "HOME_PRESENT"
	HOME_PERMANENT AddressType = "HOME_PERMANENT"

	WORK                     PhoneType = "WORK"
	PERSONAL_PHONE           PhoneType = "PERSONAL"
	HOME                     PhoneType = "HOME"
	PHONE_TYPE_NOT_SPECIFIED PhoneType = "NOT_SPECIFIED"

	ROUTING_NUMBER           RoutingType = "ROUTING_NUMBER"
	IBAN                     RoutingType = "IBAN"
	SWIFT_CODE               RoutingType = "SWIFT_CODE"
	IFSC_CODE                RoutingType = "IFSC_CODE"
	BANK_IDENTIFICATION_CODE RoutingType = "BANK_IDENTIFICATION_CODE"
	BRANCH_CODE              RoutingType = "BRANCH_CODE"

	SAVINGS               AccountType = "SAVINGS"
	CURRENT               AccountType = "CURRENT"
	CHECKING              AccountType = "CHECKING"
	ACCOUNT_NOT_SPECIFIED AccountType = "NOT_SPECIFIED"

	ALL     PayType = "ALL"
	REGULAR PayType = "REGULAR"
	EXPENSE PayType = "EXPENSE"
	BONUS   PayType = "BONUS"

	NATIONAL_ID          IDType = "NATIONAL_ID"
	LICENSE              IDType = "LICENSE"
	PASSPORT             IDType = "PASSPORT"
	OTHERS               IDType = "OTHERS"
	IDTYPE_NOT_SPECIFIED IDType = "NOT_SPECIFIED"
)
