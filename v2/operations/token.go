package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"time"
)

// Example:
//
//	token, create := &omise.Token{}, &CreateToken{
//	}
//	if e := client.Do(token, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("token: %#v\n", token)
//
type CreateToken struct {
	Card *Card `json:"card,omitempty"`
}

// Card
type Card struct {
	City string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	ExpirationMonth time.Month `json:"expiration_month,omitempty"`
	ExpirationYear int `json:"expiration_year,omitempty"`
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	SecurityCode string `json:"security_code,omitempty"`
	State string `json:"state,omitempty"`
	Street1 string `json:"street1,omitempty"`
	Street2 string `json:"street2,omitempty"`
}

// Describe
func (req *CreateToken) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.Vault,
		Method:      "POST",
		Path:        "/tokens",
		ContentType: "application/json",
	}
}

// Example:
//
//	token, retrieve := &omise.Token{}, &RetrieveToken{
//		TokenID: "tokn_5isdn6jd9c"
//	}
//	if e := client.Do(token, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("token: %#v\n", token)
//
type RetrieveToken struct {
	TokenID string `json:"-"`
}

// Describe
func (req *RetrieveToken) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.Vault,
		Method:      "GET",
		Path:        "/tokens/" + req.TokenID,
		ContentType: "application/json",
	}
}

