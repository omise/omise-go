package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Card represents Card object.
// See https://www.omise.co/cards-api for more information.
type Card struct {
	Base
	Bank string `json:"bank"`
	Brand string `json:"brand"`
	City *string `json:"city"`
	Country string `json:"country"`
	ExpirationMonth time.Month `json:"expiration_month"`
	ExpirationYear int `json:"expiration_year"`
	Financing string `json:"financing"`
	Fingerprint string `json:"fingerprint"`
	FirstDigits *string `json:"first_digits"`
	LastDigits string `json:"last_digits"`
	Location string `json:"location"`
	Name string `json:"name"`
	PhoneNumber *string `json:"phone_number"`
	PostalCode *string `json:"postal_code"`
	SecurityCodeCheck bool `json:"security_code_check"`
	State *string `json:"state"`
	Street1 *string `json:"street1"`
	Street2 *string `json:"street2"`
}

// CardService represents resource service.
type CardService struct {
	*Client
}

// Card defines resource service.
func (c *Client) Card() *CardService {
	return &CardService{c}
}

// Destroy destroys card
//
// Example:
//
//	card, destroy := &omise.Card{}, &DestroyCard{
//		CustomerID: "cust_8b6jt4hr5i"
//		CardID: "card_8b63gq9c"
//	}
//	if e := client.Do(card, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("card: %#v\n", card)
//
func (s *CardService) Destroy(ctx context.Context, params *DestroyCardParams) (*Card, error) {
	result := &Card{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DestroyCardParams params object.
type DestroyCardParams struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

// Describe describes structure of request
func (req *DestroyCardParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

// Retrieve retrieves card
//
// Example:
//
//	card, retrieve := &omise.Card{}, &RetrieveCard{
//		CustomerID: "cust_8b6jt4hr5i"
//		CardID: "card_8b63gq9c"
//	}
//	if e := client.Do(card, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("card: %#v\n", card)
//
func (s *CardService) Retrieve(ctx context.Context, params *RetrieveCardParams) (*Card, error) {
	result := &Card{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveCardParams params object.
type RetrieveCardParams struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveCardParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

// Update updates card
//
// Example:
//
//	card, update := &omise.Card{}, &UpdateCard{
//		CustomerID: "cust_8b6jt4hr5i"
//		CardID: "card_8b63gq9c"
//	}
//	if e := client.Do(card, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("card: %#v\n", card)
//
func (s *CardService) Update(ctx context.Context, params *UpdateCardParams) (*Card, error) {
	result := &Card{}
	err := s.Do(ctx, result, params)

	return result, err
}

// UpdateCardParams params object.
type UpdateCardParams struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
	City string `json:"city,omitempty"`
	ExpirationMonth time.Month `json:"expiration_month,omitempty"`
	ExpirationYear int `json:"expiration_year,omitempty"`
	Name string `json:"name,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

// Describe describes structure of request
func (req *UpdateCardParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

// List lists cards
//
// Example:
//
//	card, list := &omise.CardList{}, &ListCards{
//		CustomerID: "cust_8b6jt4hr5i"
//	}
//	if e := client.Do(card, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("card: %#v\n", card)
//
func (s *CardService) List(ctx context.Context, params *ListCardsParams) (*CardList, error) {
	result := &CardList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListCardsParams params object.
type ListCardsParams struct {
	CustomerID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListCardsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/cards",
		ContentType: "application/json",
	}
}

