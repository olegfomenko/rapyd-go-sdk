package payments

import (
	"github.com/olegfomenko/rapyd-go-sdk/resources"
)

const SPEIPaymentType = "mx_spei_bank"

func SPEI(amount string, description string, expiration int64, metadata interface{}) *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{
			"amount":      amount,
			"description": description,
			"expiration":  expiration,
			"metadata":    metadata,
		},
		Type: SPEIPaymentType,
	}
}
