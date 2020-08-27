package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"time"
)

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
type DestroyCard struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

// Describe
func (req *DestroyCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

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
type RetrieveCard struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

// Describe
func (req *RetrieveCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

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
type UpdateCard struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
	City string `json:"city,omitempty"`
	ExpirationMonth time.Month `json:"expiration_month,omitempty"`
	ExpirationYear int `json:"expiration_year,omitempty"`
	Name string `json:"name,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

// Describe
func (req *UpdateCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

// Example:
//
//	card, list := &omise.Card{}, &ListCards{
//		CustomerID: "cust_8b6jt4hr5i"
//	}
//	if e := client.Do(card, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("card: %#v\n", card)
//
type ListCards struct {
	CustomerID string `json:"-"`
	List
}

// Describe
func (req *ListCards) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/cards",
		ContentType: "application/json",
	}
}

