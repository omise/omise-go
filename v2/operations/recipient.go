package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"github.com/omise/omise-go/v2"
)

// Example:
//
//	charge, update := &omise.Recipient{}, &DestroyRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyRecipient struct {
	RecipientID string `json:"-"`
}

func (req *DestroyRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &RetrieveRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveRecipient struct {
	RecipientID string `json:"-"`
}

func (req *RetrieveRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &UpdateRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateRecipient struct {
	RecipientID string `json:"-"`
	BankAccount *omise.BankAccount `json:"bank_account"`
	Description string `json:"description"`
	Email string `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
	Name string `json:"name"`
	TaxID string `json:"tax_id"`
	Type *omise.RecipientType `json:"type"`
}

func (req *UpdateRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &SearchRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchRecipient struct {
}

func (req *SearchRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &ListRecipients{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListRecipients struct {
}

func (req *ListRecipients) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &CreateRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateRecipient struct {
	BankAccount *omise.BankAccount `json:"bank_account"`
	Description string `json:"description"`
	Email string `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
	Name string `json:"name"`
	TaxID string `json:"tax_id"`
	Type *omise.RecipientType `json:"type"`
}

func (req *CreateRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/recipients",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &VerifyRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type VerifyRecipient struct {
	RecipientID string `json:"-"`
}

func (req *VerifyRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/recipients/" + req.RecipientID + "/verify",
		ContentType: "application/json",
	}
}

