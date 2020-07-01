package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Token represents Token object.
// See https://www.omise.co/tokens-api for more information.
type Token struct {
	Base
	Card *Card `json:"card"`
	ChargeStatus ChargeStatus `json:"charge_status"`
	Location string `json:"location"`
	Used bool `json:"used"`
}

// TokenService represents resource service.
type TokenService struct {
	*Client
}

// Token defines resource service.
func (c *Client) Token() *TokenService {
	return &TokenService{c}
}

// Create creates token
//
// Example:
//
//	token, err := client.Token().Create(ctx, &CreateTokenParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("token: %#v\n", token)
//
func (s *TokenService) Create(ctx context.Context, params *CreateTokenParams) (*Token, error) {
	result := &Token{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateTokenParams params object.
type CreateTokenParams struct {
	Card *CardParams `json:"card,omitempty"`
}

// CardParams params object.
type CardParams struct {
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

// Describe describes structure of request
func (req *CreateTokenParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.Vault,
		Method:      "POST",
		Path:        "/tokens",
		ContentType: "application/json",
	}
}

// Retrieve retrieves token
//
// Example:
//
//	token, err := client.Token().Retrieve(ctx, &RetrieveTokenParams{
//		TokenID: "tokn_5isdn6jd9c"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("token: %#v\n", token)
//
func (s *TokenService) Retrieve(ctx context.Context, params *RetrieveTokenParams) (*Token, error) {
	result := &Token{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveTokenParams params object.
type RetrieveTokenParams struct {
	TokenID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveTokenParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.Vault,
		Method:      "GET",
		Path:        "/tokens/" + req.TokenID,
		ContentType: "application/json",
	}
}

