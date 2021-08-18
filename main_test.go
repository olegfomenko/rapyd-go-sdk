package rapyd

import (
	resources2 "github.com/olegfomenko/rapyd-go-sdk/resources"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

const (
	accessKey = ""
	secretKey = ""
	endpoint  = "https://sandboxapi.rapyd.net"
)

func TestClient_CreateWallet(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient([]byte(accessKey), []byte(secretKey), addr, http.DefaultClient)

	_, err = rapyd.CreateWallet(resources2.Wallet{
		FirstName:   "Oleg",
		LastName:    "Fomenko",
		Email:       "127@rapyd.net",
		Reference:   "127-Oleg-20072021",
		PhoneNumber: "+14155551117",
		Type:        resources2.PersonWalletType,
		Contact: resources2.Contact{
			PhoneNumber: "+14155551117",
			Email:       "127@rapyd.net",
			FirstName:   "Oleg",
			LastName:    "Fomenko",
			ContactType: resources2.PersonalContactType,
			Address: resources2.Address{
				Name:    "Oleg Fomenko",
				Line1:   "124 Main Street",
				Line2:   "",
				Line3:   "",
				City:    "Anytown",
				State:   "NY",
				Country: "US",
				Zip:     "12345",
				Phone:   "+14155551117",
			},
			IdentificationType:   "PA",
			IdentificationNumber: "1234567891",
			Birth:                "11/22/2000",
			Country:              "US",
			Nationality:          "FR",
		},
	})

	assert.NoError(t, err)
}

func TestClient_CreateCustomer(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient([]byte(accessKey), []byte(secretKey), addr, http.DefaultClient)
	_, err = rapyd.CreateCustomer(resources2.Customer{
		Name:  "Oleg Fomenko",
		Email: "111115@rapyd.net",
		Phone: "+3800661111115",
		PaymentMethod: resources2.PaymentMethod{
			Fields: map[string]interface{}{
				"proof_of_authorization": false,
				"routing_number":         "111111111",
				"payment_purpose":        "111111",
				"account_number":         "111111",
			},
			Type: "us_ach_bank",
		},
	})

	assert.NoError(t, err)
}

func TestClient_CreatePayment(t *testing.T) {
	addr, err := url.Parse(endpoint)
	assert.NoError(t, err)

	rapyd := NewClient([]byte(accessKey), []byte(secretKey), addr, http.DefaultClient)
	_, err = rapyd.CreatePayment(resources2.CreatePayment{
		Amount:   "100.23",
		Currency: "USD",
		PaymentMethod: &resources2.PaymentMethod{
			Fields: map[string]interface{}{
				"proof_of_authorization": false,
				"routing_number":         "111111111",
				"payment_purpose":        "111111",
				"account_number":         "111111",
			},
			Type: "us_ach_bank",
		},
		EWallets: []resources2.EWallet{
			{
				Wallet:     "ewallet_8d3fdd0929856f5a30ec2933f4bd6cf1",
				Percentage: 100,
			},
		},
	})

	assert.NoError(t, err)
}
