package resources

type Sender struct {
	CompanyName             *string `json:"company_name,omitempty"`
	Country                 string  `json:"country"`
	Currency                string  `json:"currency"`
	EntityType              string  `json:"entity_type"`
	FirstName               string  `json:"first_name"`
	IdentificationType      string  `json:"identification_type"`
	IdentificationValue     string  `json:"identification_value"`
	LastName                string  `json:"last_name"`
	PhoneNumber             string  `json:"phone_number"`
	Occupation              string  `json:"occupation"`
	SourceOfIncome          string  `json:"source_of_income"`
	DateOfBirth             string  `json:"date_of_birth"`
	Address                 string  `json:"address"`
	City                    string  `json:"city"`
	State                   string  `json:"state"`
	PostCode                string  `json:"postcode"`
	PurposeCode             string  `json:"purpose_code"`
	BeneficiaryRelationship string  `json:"beneficiary_relationship"`
}

type SenderResponse struct {
	Data Data `json:"data"`
}
