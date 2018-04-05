package operations

import (
	"time"

	"github.com/omise/omise-go/internal"
)

// Example:
//
//	token, create := &omise.Token{}, &CreateToken{
//		Name:            "Chakrit Wichian",
//		Number:          "4242424242424242",
//		ExpirationMonth: 2,
//		ExpirationYear:  time.Now().AddDate(1, 0, 0).Year(),
//
//		SecurityCode: "123",
//		City:         "Bangkok",
//		PostalCode:   "10240",
//	}
//	if e := client.Do(token, CreateTokenOp); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created token:", token.ID)
//
type CreateToken struct {
	Name            string     `query:"card[name]"`
	Number          string     `query:"card[number]"`
	ExpirationMonth time.Month `query:"card[expiration_month]"`
	ExpirationYear  int        `query:"card[expiration_year]"`

	SecurityCode string `query:"card[security_code]"`
	City         string `query:"card[city]"`
	PostalCode   string `query:"card[postal_code]"`
}

func (req *CreateToken) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.Vault,
		Method:   "POST",
		Path:     "/tokens",
	}
}

// Example:
//
//	token, retrieve := &omise.Token{}, &RetrieveToken{"tok_789"}
//	if e := client.Do(token, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("token: %#v\n", token)
//
type RetrieveToken struct {
	ID string `query:"-"`
}

func (token *RetrieveToken) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.Vault,
		Method:   "GET",
		Path:     "/tokens/" + token.ID,
	}
}
