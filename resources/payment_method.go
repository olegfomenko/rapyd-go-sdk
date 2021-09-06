package resources

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Regex       string `json:"regex"`
	Description string `json:"description"`
	Required    bool   `json:"is_required"`
	Updatable   bool   `json:"is_updatable"`
}

type PaymentMethodRequiredFields struct {
	Type                 string  `json:"type"`
	Fields               []Field `json:"fields"`
	MethodOptions        []Field `json:"payment_method_options"`
	PaymentOptions       []Field `json:"payment_options"`
	MinExpirationSeconds int64   `json:"minimum_expiration_seconds"`
	MaxExpirationSeconds int64   `json:"maximum_expiration_seconds"`
}

type PaymentMethodRequiredFieldsResponse struct {
	Data PaymentMethodRequiredFields `json:"data"`
}
