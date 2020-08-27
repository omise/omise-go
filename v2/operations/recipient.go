package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"github.com/omise/omise-go/v2"
)

// Example:
//
//	recipient, destroy := &omise.Recipient{}, &DestroyRecipient{
//		RecipientID: "recip_3gqd8b4h1e"
//	}
//	if e := client.Do(recipient, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient: %#v\n", recipient)
//
type DestroyRecipient struct {
	RecipientID string `json:"-"`
}

// Describe
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
//	recipient, retrieve := &omise.Recipient{}, &RetrieveRecipient{
//		RecipientID: "recip_3gqd8b4h1e"
//	}
//	if e := client.Do(recipient, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient: %#v\n", recipient)
//
type RetrieveRecipient struct {
	RecipientID string `json:"-"`
}

// Describe
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
//	recipient, update := &omise.Recipient{}, &UpdateRecipient{
//		RecipientID: "recip_3gqd8b4h1e"
//	}
//	if e := client.Do(recipient, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient: %#v\n", recipient)
//
type UpdateRecipient struct {
	RecipientID string `json:"-"`
	BankAccount *BankAccount `json:"bank_account,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name,omitempty"`
	TaxID string `json:"tax_id,omitempty"`
	Type omise.RecipientType `json:"type,omitempty"`
}

// BankAccount
type BankAccount struct {
	BankCode string `json:"bank_code,omitempty"`
	BranchCode string `json:"branch_code,omitempty"`
	Brand string `json:"brand,omitempty"`
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	Type omise.BankAccountType `json:"type,omitempty"`
}

// Describe
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
//	recipient, list := &omise.Recipient{}, &ListRecipients{
//	}
//	if e := client.Do(recipient, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient: %#v\n", recipient)
//
type ListRecipients struct {
	List
}

// Describe
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
//	recipient, create := &omise.Recipient{}, &CreateRecipient{
//	}
//	if e := client.Do(recipient, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient: %#v\n", recipient)
//
type CreateRecipient struct {
	BankAccount *BankAccount `json:"bank_account,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name,omitempty"`
	TaxID string `json:"tax_id,omitempty"`
	Type omise.RecipientType `json:"type,omitempty"`
}

// Describe
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
//	schedule, listSchedules := &omise.Schedule{}, &ListRecipientSchedules{
//		RecipientID: "recip_3gqd8b4h1e"
//	}
//	if e := client.Do(schedule, listSchedules); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
type ListRecipientSchedules struct {
	RecipientID string `json:"-"`
	List
}

// Describe
func (req *ListRecipientSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients/" + req.RecipientID + "/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	recipient, verify := &omise.Recipient{}, &VerifyRecipient{
//		RecipientID: "recip_3gqd8b4h1e"
//	}
//	if e := client.Do(recipient, verify); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient: %#v\n", recipient)
//
type VerifyRecipient struct {
	RecipientID string `json:"-"`
}

// Describe
func (req *VerifyRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/recipients/" + req.RecipientID + "/verify",
		ContentType: "application/json",
	}
}

