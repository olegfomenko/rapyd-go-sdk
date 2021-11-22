package resources

type CreatePayout struct {
	Beneficiary           Beneficiary `json:"beneficiary"`
	BeneficiaryCountry    string      `json:"beneficiary_country,omitempty"`
	BeneficiaryEntityType string      `json:"beneficiary_entity_type"`
	ConfirmAutomatically  *bool       `json:"confirm_automatically,omitempty"`
	Description           *string     `json:"description,omitempty"`
	Expiration            *string     `json:"expiration,omitempty"`
	EWallet               string      `json:"ewallet,omitempty"`
	PayoutAmount          float64     `json:"payout_amount"`
	PayoutCurrency        string      `json:"payout_currency"`
	PayoutMethodType      string      `json:"payout_method_type"`
	Sender                Sender      `json:"sender"`
	SenderAmount          *float64    `json:"sender_amount,omitempty"`
	SenderCountry         string      `json:"sender_country"`
	SenderCurrency        string      `json:"sender_currency"`
	SenderEntityType      string      `json:"sender_entity_type"`
	StatementDescriptor   *string     `json:"statement_descriptor,omitempty"`
	CardNumber            string      `json:"card_number"`
}

type CreatePayoutResponse struct {
	Data Data `json:"data"`
}
