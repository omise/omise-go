package operations

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	recipients, list := &omise.RecipientList{}, &ListRecipients{
//		List{Offset: 100, Limit: 20},
//	}
//	if e := client.Do(recipients, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("recipients #100-#120:", recipients.Data)
//
type ListRecipients struct {
	List
}

func (req *ListRecipients) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/recipients",
	}
}

// Example:
//
//	bankAccount := &omise.BankAccount{
//		Brand:  "bbl",
//		Number: "1234567890",
//		Name:   "Joe Example",
//	}
//
//	jun, create := &omise.Recipient{}, &CreateRecipient{
//		Name:        "Jun Hasegawa",
//		Email:       "jun@omise.co",
//		Description: "Owns Omise",
//		Type:        omise.Individual,
//		BankAccount: bankAccount,
//	}
//	if e := client.Do(jun, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created recipient:", jun.ID)
//
type CreateRecipient struct {
	Name        string
	Email       string
	Description string
	Type        omise.RecipientType
	TaxID       string             `query:"tax_id"`
	BankAccount *omise.BankAccount `query:"bank_account"`
}

func (req *CreateRecipient) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "POST",
		Path:     "/recipients",
	}
}

// Example:
//
//	recp, retrieve := &omise.Recipient{}, &RetrieveRecipient{"recp_123"}
//	if e := client.Do(recp, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient #123: %#v\n", recp)
//
type RetrieveRecipient struct {
	RecipientID string `query:"-"`
}

func (req *RetrieveRecipient) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/recipients/" + req.RecipientID,
	}
}

// Example:
//
//	jones, update := &omise.Recipient{}, &UpdateRecipient{
//		RecipientID: "recp_123",
//		Description: "I'm JONES now.",
//	}
//	if e := client.Do(jones, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("jones: %#v\n", jones)
//
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
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "PATCH",
		Path:     "/recipients/" + req.RecipientID,
	}
}

// Example:
//
//	del, destroy := &omise.Deletion{}, &DestroyRecipient{"recp-123"}
//	if e := client.Do(del, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("destroyed recipient:", del.ID)
//
type DestroyRecipient struct {
	RecipientID string `query:"-"`
}

func (req *DestroyRecipient) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "DELETE",
		Path:     "/recipients/" + req.RecipientID,
	}
}
