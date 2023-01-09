package wallet

type Params struct {
	Contact            interface{} `json:"contact"`
	EWalletReferenceId string      `json:"ewallet_reference_id"`
	FirstName          string      `json:"first_name"`
	LastName           string      `json:"last_name"`
	Metadata           interface{} `json:"metadata"`
	Type               string      `json:"type"` // personal, client, or business
}

type PersonalContact struct {
	Address     interface{} `json:"address"`
	ContactType string      `json:"contact_type"` // expected type: 'personal'
	Country     string      `json:"country"`      // UPPERCASE two letter ISO-3166-1 Alpha-2 code, example: Argentina == AR, Aruba == AW
	// func main.GetCountries can be called to see Rapyd's list
	// country list:
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
	DateOfBirth          string `json:"date_of_birth"` // expected format: MM/DD/YY
	Email                string `json:"email"`
	HouseType            string `json:"house_type"` //expected:  lease, live_with_family, own, owner, month_to_month, housing_project
	IdentificationNumber string `json:"identification_number"`
	IdentificationType   string `json:"identification_type"` // GET list of docs by country from /v1/identities/types?country
	// country param is 2 char alphanumeric iso 3166-1 Alpha-2. See above
	LastName          string      `json:"last_name"`
	MaritalStatus     string      `json:"marital_status"` // expected: married, single, divorced, widowed, cohabiting, not_applicable
	Metadata          interface{} `json:"metadata"`       // json obj determiend by the client
	MiddleName        string      `json:"middle_name"`
	MothersName       string      `json:"mothers_name"`
	Nationality       string      `json:"nationality"`  // see country code stuff above ^^
	PhoneNumber       string      `json:"phone_number"` // phone number in E.164 format: +16175551212
	SecondLastName    string      `json:"second_last_name"`
	SendNotifications bool        `json:"send_notifications"` // defaults to false
}
