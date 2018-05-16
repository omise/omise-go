package operations

import (
	"encoding/json"
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
	Name            string     `json:"name"`
	Number          string     `json:"number"`
	ExpirationMonth time.Month `json:"expiration_month"`
	ExpirationYear  int        `json:"expiration_year"`

	SecurityCode string `json:"security_code"`
	City         string `json:"city,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
}

func (req *CreateToken) MarshalJSON() ([]byte, error) {
	type CardToken CreateToken
	params := struct {
		Card *CardToken `json:"card"`
	}{
		Card: (*CardToken)(req),
	}
	return json.Marshal(params)
}

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
//	token, retrieve := &omise.Token{}, &RetrieveToken{"tok_789"}
//	if e := client.Do(token, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("token: %#v\n", token)
//
type RetrieveToken struct {
	ID string `json:"-"`
}

func (token *RetrieveToken) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.Vault,
		Method:      "GET",
		Path:        "/tokens/" + token.ID,
		ContentType: "application/json",
	}
}
