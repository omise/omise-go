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

func (req *CreateToken) Op() *internal.Op {
	return &internal.Op{internal.Vault, "POST", "/tokens", nil}
}

// Example:
//
//	token, retreive := &omise.Token{}, &RetreiveToken{"tok_789"}
//	if e := client.Do(token, retreive); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("token: %#v\n", token)
//
type RetreiveToken struct {
	ID string `query:"-"`
}

func (token *RetreiveToken) Op() *internal.Op {
	return &internal.Op{internal.Vault, "GET", "/tokens/" + token.ID, nil}
}
