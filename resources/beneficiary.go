package resources

type Beneficiary struct {
	Category                string  `json:"category"`
	CompanyName             *string `json:"company_name,omitempty"`
	Country                 string  `json:"country"`
	Currency                string  `json:"currency"`
	DefaultPayoutMethodType *string `json:"default_payout_method_type,omitempty"`
	EntityType              string  `json:"entity_type"`
	FirstName               string  `json:"first_name"`
	IdentificationType      string  `json:"identification_type"`
	IdentificationValue     string  `json:"identification_value"`
	LastName                string  `json:"last_name"`
	AccountNumber           string  `json:"account_number"`
	PaymentType             string  `json:"payment_type"`
	Address                 string  `json:"address"`
	City                    string  `json:"city"`
	State                   string  `json:"state"`
	PostCode                string  `json:"postcode"`
	CardNumber              string  `json:"card_number"`
	CardExpirationMonth     string  `json:"card_expiration_month"`
	CardExpirationYear      string  `json:"card_expiration_year"`
	AchCode                 string  `json:"aba"`
	BankName                string  `json:"bank_name"`
	BicSwift                string  `json:"bic_swift"`
}

type BeneficiaryResponse struct {
	Data Data `json:"data"`
}
