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

func (req *ListCards) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/customers/" + req.CustomerID + "/cards",
	}
}

// TODO: Cards can be created only by updating an existing customer and the result will be
//   a Customer instance instead of a Card. So the API's a bit confusing here.
// type CreateCard struct {
// 	CustomerID string `query:"-"`
// 	Card       string
// }
//
// func (req *CreateCard) Description() *internal.Description {
// 	return &internal.Description{internal.API, "PATCH", "/customers/" + req.CustomerID, nil}
// }

// Example:
//
//	card, retrieve := &omise.Card{}, &RetrieveCard{
//		CustomerID: "cust_888",
//		CardID:     "card_999",
//	}
//	if e := client.Do(card, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("the card: %#v\n", card)
//
type RetrieveCard struct {
	CustomerID string `query:"-"`
	CardID     string `query:"-"`
}

func (req *RetrieveCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/customers/" + req.CustomerID + "/cards/" + req.CardID,
	}
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

func (req *UpdateCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "PATCH",
		Path:     "/customers/" + req.CustomerID + "/cards/" + req.CardID,
	}
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

func (req *DestroyCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "DELETE",
		Path:     "/customers/" + req.CustomerID + "/cards/" + req.CardID,
	}
}
