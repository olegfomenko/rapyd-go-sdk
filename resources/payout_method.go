package resources

type PayoutRequiredFields struct {
	SenderCurrency            string                 `json:"sender_currency"`
	SenderCountry             string                 `json:"sender_country"`
	SenderEntityType          string                 `json:"sender_entity_type"`
	BeneficiaryCountry        string                 `json:"beneficiary_country"`
	PayoutCurrency            string                 `json:"payout_currency"`
	BeneficiaryEntityType     string                 `json:"beneficiary_entity_type"`
	IsCancelable              bool                   `json:"is_cancelable"`
	IsLocationSpecific        bool                   `json:"is_location_specific"`
	IsExpirable               bool                   `json:"is_expirable"`
	IsOnline                  bool                   `json:"is_online"`
	Image                     string                 `json:"image"`
	Status                    bool                   `json:"status"`
	BeneficiaryRequiredFields BeneficiarySenderField `json:"beneficiary_required_fields"`
	SenderRequiredFields      BeneficiarySenderField `json:"sender_required_fields"`
}

type BeneficiarySenderField struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Regex       string `json:"regex"`
	Description string `json:"description"`
}

type PayoutRequiredFieldsResponse struct {
	Data PayoutRequiredFields `json:"data"`
}
