package japan

import "github.com/olegfomenko/rapyd-go-sdk/resources"

const JPBankTransferPaymentType = "jp_banktransfer_bank"

func JPBank() *resources.PaymentMethod {
	return &resources.PaymentMethod{
		Fields: map[string]interface{}{},
		Type:   JPBankTransferPaymentType,
	}
}
