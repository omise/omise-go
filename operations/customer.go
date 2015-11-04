package operations

import (
	"github.com/omise/omise-go/internal"
)

type ListCustomers struct {
	List
}

func (req *ListCustomers) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers", nil}
}

type CreateCustomer struct {
	Email       string
	Description string
	Card        string
}

func (req *CreateCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/customers", nil}
}

type RetreiveCustomer struct {
	CustomerID string `query:"-"`
}

func (req *RetreiveCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers/" + req.CustomerID, nil}
}

type UpdateCustomer struct {
	CustomerID  string `query:"-"`
	Email       string
	Description string
	Card        string
}

func (req *UpdateCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/customers/" + req.CustomerID, nil}
}

type DestroyCustomer struct {
	CustomerID string `query:"-"`
}

func (req *DestroyCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "DELETE", "/customers/" + req.CustomerID, nil}
}
