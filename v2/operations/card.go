package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Card{}, &DestroyCard{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyCard struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

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
//	charge, update := &omise.Card{}, &RetrieveCard{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveCard struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

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
//	charge, update := &omise.Card{}, &UpdateCard{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateCard struct {
	CustomerID string `json:"-"`
	CardID string `json:"-"`
	City string `json:"city"`
	ExpirationMonth int `json:"expiration_month"`
	ExpirationYear int `json:"expiration_year"`
	Name string `json:"name"`
	PostalCode string `json:"postal_code"`
}

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
//	charge, update := &omise.Card{}, &ListCards{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListCards struct {
	CustomerID string `json:"-"`
}

func (req *ListCards) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/cards",
		ContentType: "application/json",
	}
}

