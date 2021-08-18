package resources

type Customer struct {
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone_number"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}

type CustomerResponse struct {
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}
