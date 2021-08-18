package resources

type ContactType string
type WalletType string

const (
	PersonalContactType ContactType = "personal"
	PersonWalletType    WalletType  = "person"
	CompanyWalletType   WalletType  = "company"
)

type Address struct {
	Name    string `json:"name"`
	Line1   string `json:"line_1"`
	Line2   string `json:"line_2"`
	Line3   string `json:"line_3"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
	Phone   string `json:"phone_number"`
}

type Contact struct {
	PhoneNumber          string      `json:"phone_number"`
	Email                string      `json:"email"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	ContactType          ContactType `json:"contact_type"`
	Address              Address     `json:"address"`
	IdentificationType   string      `json:"identification_type"`
	IdentificationNumber string      `json:"identification_number"`
	Birth                string      `json:"date_of_birth"`
	Country              string      `json:"country"`
	Nationality          string      `json:"nationality"`
}

type Wallet struct {
	Id          string     `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Reference   string     `json:"ewallet_reference_id"`
	PhoneNumber string     `json:"phone_number"`
	Type        WalletType `json:"type"`
	Contact     Contact    `json:"contact"`
}

type WalletResponse struct {
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}
