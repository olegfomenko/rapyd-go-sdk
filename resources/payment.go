package resources

const (
	ProofOfAuthorizationField = "proof_of_authorization"
	RoutingNumberField        = "routing_number"
	PaymentPurposeField       = "payment_purpose"
	AccountNumberField        = "account_number"
	DefaultPaymentType        = "us_ach_bank"
)

type PaymentMethod struct {
	Fields map[string]interface{} `json:"fields"`
	Type   string                 `json:"type"`
}

type EWallet struct {
	Wallet     string `json:"ewallet"`
	Percentage int32  `json:"percentage"`
}

type CreatePayment struct {
	Amount        string         `json:"amount"`
	Currency      string         `json:"currency"`
	Customer      *string        `json:"customer,omitempty"`
	EWallets      []EWallet      `json:"ewallets"`
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

type CreatePaymentResponse struct {
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}
