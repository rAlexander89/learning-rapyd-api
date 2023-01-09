package wallet

// // // NOTES
// Personal DOB vs Business Established Date:
// The Personal DOB field is expectd to be formatted as MM/DD/YY
// Business Established date is expecated to be YYY-MM-DD

// Best Practice:
// 1. For Personal contact, include
// first_name,
// last_name,
// date_of_birth,
// country,
// address.name,
// address.line_1,
//
// These fields are used in regularly on many wallet operations.
// // //

// // // Start: Two Contact types
// 1. Personal
// 2. Business
//

// 1. Personal Contact
type Contact struct {
	ID                string      `json:"id"` //prefixed: 'cont_'
	Address           interface{} `json:"address"`
	BusinessDetails   interface{} `json:"business_details"` // required when Wallet type is 'company'
	ComplianceProfile int         `json:"compliance_profile"`
	ContactType       string      `json:"contact_type"` // personal (individual customer) or business (business customer)
	Country           string      `json:"country"`      // UPPERCASE two letter ISO-3166-1 Alpha-2 code, example: Argentina == AR, Aruba == AW
	// func main.GetCountries can be called to see Rapyd's list
	// country list:
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
	CreatedAt            int         `json:"created_at"`    // unix time. Is int the best data type for this?
	DateOfBirth          string      `json:"date_of_birth"` // expected format: MM/DD/YYY
	Email                string      `json:"email"`
	EWallet              string      `json:"ewallet"`
	FirstName            string      `json:"first_name"`
	Gender               string      `json:"gender"`
	HouseType            string      `json:"house_type"`
	IdentificationNumber string      `json:"identification_number"`
	IdentificationType   string      `json:"identification_type"`
	IssuedCardData       interface{} `json:"issued_card_data"`
	LastName             string      `json:"last_name"`
	MaritalStatus        string      `json:"marital_status"`
	Metadata             interface{} `json:"metadata"`
	MiddleName           string      `json:"middle_name"`
	MothersName          string      `json:"mothers_name"`
	Nationality          string      `json:"nationality"` // UPPERCASE two letter ISO-3166-1 Alpha-2 code, example: Argentina == AR, Aruba == AW
	// func main.GetCountries can be called to see Rapyd's list
	// country list:
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
	PhoneNumner       string `json:"phone_number"`
	SecondLastName    string `json:"second_last_name"`
	SendNotifications bool   `json:"send_notifications"`  // defaults to false
	VerificationStaus string `json:"verification_status"` // KYC status: 'not verified', 'KYCd'
}

// 2. Business Contact
type BusinessDetails struct {
	ID                  string      `json:"id"` //prefixed: 'busi_'
	Address             interface{} `json:"address"`
	AnnualRevenue       float64     `json:"annual_revenue"`
	CnaeCode            string      `json:"cnae_code"` // Business Activity code. AlphaNumeric. max len == 7
	CreatedAt           int         `json:"created_at"`
	EntityType          string      `json:"entity_type"`        //Business type: sole_prop, partnership, company, government, charity, NPO, association, trust
	EstablishmentDate   string      `json:"establishment_date"` // expected format: YYY-MM-DD
	IndustryCategory    string      `json:"industry_category"`  // Required. Alphanumeric with no special cahrs
	IndustrySubCategory string      `json:"industry_sub_category"`
	LegalEntityType     string      `json:"legal_entity_type"`
	Name                string      `json:"name"`
	RegistrationNumber  string      `json:"registration_number"`
}

// // // End: Two Contact types
