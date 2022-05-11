package resources

type Customer struct {
	Addresses            []Address     `json:"addresses"`
	BusinessVatID        string        `json:"business_vat_id"`
	Coupon               string        `json:"coupon"`
	DefaultPaymentMethod string        `json:"default_payment_method"`
	Description          string        `json:"description"`
	EWallet              string        `json:"e_wallet"`
	Name                 string        `json:"name"`
	Email                string        `json:"email"`
	Phone                string        `json:"phone_number"`
	PaymentMethod        PaymentMethod `json:"payment_method"`
}

type CustomerResponse struct {
	Data Data `json:"data"`
}
