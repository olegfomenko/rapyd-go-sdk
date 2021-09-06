package philippines

import "github.com/olegfomenko/rapyd-go-sdk/resources"

const BancnetPaymentType = "ph_bancnet_bank"

func BancnetBank(description string) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"description": description,
		},
		Type: BancnetPaymentType,
	}
}
