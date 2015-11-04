package operations

import (
	"time"

	"github.com/omise/omise-go/internal"
)

type ListCards struct {
	List
	CustomerID string `query:"-"`
}

func (req *ListCards) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers/" + req.CustomerID + "/cards", nil}
}

type RetreiveCard struct {
	CustomerID string `query:"-"`
	CardID     string `query:"-"`
}

func (req *RetreiveCard) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers/" + req.CustomerID + "/cards/" + req.CardID, nil}
}

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

type DestroyCard struct {
	CustomerID string `query:"-"`
	CardID     string `query:"-"`
}

func (req *DestroyCard) Op() *internal.Op {
	return &internal.Op{internal.API, "DELETE", "/customers/" + req.CustomerID + "/cards/" + req.CardID, nil}
}
