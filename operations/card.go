package operations

import (
	"time"

	"github.com/omise/omise-go/internal"
)

// Example:
//
//	cards, list := &omise.CardList{}, &ListCards{
//		CustomerID: "cust_333",
//	}
//	if e := client.Do(cards, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of customer's cards:", len(cards.Data))
//
type ListCards struct {
	CustomerID string `query:"-"`
	List
}

func (req *ListCards) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers/" + req.CustomerID + "/cards", nil}
}

// Example:
//
//	card, retreive := &omise.Card{}, &RetreiveCard{
//		CustomerID: "cust_888",
//		CardID:     "card_999",
//	}
//	if e := client.Do(card, retreive); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("the card: %#v\n", card)
//
type RetreiveCard struct {
	CustomerID string `query:"-"`
	CardID     string `query:"-"`
}

func (req *RetreiveCard) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers/" + req.CustomerID + "/cards/" + req.CardID, nil}
}

// Example:
//
//	card, update := &omise.Card{}, &UpdateCard{
//		CustomerID: "cust_123",
//		CardID:     "card_123",
//		Name:       "Changed my name",
//	}
//	if e := client.Do(card, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated card: %#v\n", card)
//
type UpdateCard struct {
	CustomerID string `query:"-"`
	CardID     string `query:"-"`

	Name       string
	City       string
	PostalCode string `query:"postal_code"`

	ExpirationMonth time.Month `query:"expiration_month"`
	ExpirationYear  int        `query:"expiration_year"`
}

func (req *UpdateCard) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/customers/" + req.CustomerID + "/cards/" + req.CardID, nil}
}

// Example:
//
//	del, destroy := &omise.Deletion{}, &DestroyCard{
//		CustomerID: "cust_123",
//		CardID:     "card_456",
//	}
//	if e := client.Do(del, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("deleted:", del.ID, del.Deleted)
//
type DestroyCard struct {
	CustomerID string `query:"-"`
	CardID     string `query:"-"`
}

func (req *DestroyCard) Op() *internal.Op {
	return &internal.Op{internal.API, "DELETE", "/customers/" + req.CustomerID + "/cards/" + req.CardID, nil}
}
