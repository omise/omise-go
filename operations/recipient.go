package operations

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
)

type ListRecipients struct {
	List
}

func (req *ListRecipients) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/recipients", nil}
}

type CreateRecipient struct {
	Name        string
	Email       string
	Description string
	Type        omise.RecipientType
	TaxID       string             `query:"tax_id"`
	BankAccount *omise.BankAccount `query:"bank_account"`
}

func (req *CreateRecipient) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/recipients", nil}
}

type RetreiveRecipient struct {
	RecipientID string `query:"-"`
}

func (req *RetreiveRecipient) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/recipients/" + req.RecipientID, nil}
}

type UpdateRecipient struct {
	RecipientID string `query:"-"`
	Name        string
	Email       string
	Description string
	Type        omise.RecipientType
	TaxID       string             `query:"tax_id"`
	BankAccount *omise.BankAccount `query:"bank_account"`
}

func (req *UpdateRecipient) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/recipients/" + req.RecipientID, nil}
}

type DestroyRecipient struct {
	RecipientID string `query:"-"`
}

func (req *DestroyRecipient) Op() *internal.Op {
	return &internal.Op{internal.API, "DELETE", "/recipients/" + req.RecipientID, nil}
}
