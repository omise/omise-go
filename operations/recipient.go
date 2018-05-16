package operations

import (
	"encoding/json"

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
//      # For Japan Bank Account
//	bankAccount := &omise.BankAccount{
//		BankCode:    "0001",
//		BranchCode:  "001",
//  	        AccountType: omise.Normal,
//  	        Number:      "0000001",
//		Name:        "Joe Example",
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
	Name        string              `json:"name"`
	Email       string              `json:"email,omitempty"`
	Description string              `json:"description,omitempty"`
	Type        omise.RecipientType `json:"type"`
	TaxID       string              `json:"tax_id,omitempty"`
	BankAccount *omise.BankAccount  `json:"-"`
}

type bankAccountParams struct {
	Brand  string `json:"brand,omitempty"`
	Number string `json:"number"`
	Name   string `json:"name"`

	// for Omise Japan
	BankCode    string                `json:"bank_code,omitempty"`
	BranchCode  string                `json:"branch_code,omitempty"`
	AccountType omise.BankAccountType `json:"account_type,omitempty"`
}

func (req *CreateRecipient) MarshalJSON() ([]byte, error) {
	type Alias CreateRecipient
	params := struct {
		*Alias
		BankAccountParams bankAccountParams `json:"bank_account"`
	}{
		Alias: (*Alias)(req),
	}
	if ba := params.BankAccount; ba != nil {
		params.BankAccountParams = bankAccountParams{
			Brand:       ba.Brand,
			Number:      ba.Number,
			Name:        ba.Name,
			BankCode:    ba.BankCode,
			BranchCode:  ba.BranchCode,
			AccountType: ba.AccountType,
		}
	}
	return json.Marshal(params)
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
//	recp, retrieve := &omise.Recipient{}, &RetrieveRecipient{"recp_123"}
//	if e := client.Do(recp, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient #123: %#v\n", recp)
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
	RecipientID string              `json:"-"`
	Name        string              `json:"name"`
	Email       string              `json:"email,omitempty"`
	Description string              `json:"description,omitempty"`
	Type        omise.RecipientType `json:"type"`
	TaxID       string              `json:"tax_id,omitempty"`
	BankAccount *omise.BankAccount  `json:"-"`
}

func (req *UpdateRecipient) MarshalJSON() ([]byte, error) {
	type Alias UpdateRecipient
	params := struct {
		*Alias
		BankAccountParams bankAccountParams `json:"bank_account"`
	}{
		Alias: (*Alias)(req),
	}
	if ba := params.BankAccount; ba != nil {
		params.BankAccountParams = bankAccountParams{
			Brand:       ba.Brand,
			Number:      ba.Number,
			Name:        ba.Name,
			BankCode:    ba.BankCode,
			BranchCode:  ba.BranchCode,
			AccountType: ba.AccountType,
		}
	}
	return json.Marshal(params)
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
//	del, destroy := &omise.Deletion{}, &DestroyRecipient{"recp-123"}
//	if e := client.Do(del, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("destroyed recipient:", del.ID)
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

// ListRecipientTransferSchedules represent list recipient transfer schedules API payload
//
// Example:
//
//      var schds omise.ScheduleList
//	list := ListRecipientTransferSchedules{
//		RecipientID: "reci_123"
//		List: List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(&schds, &list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of schedules made in the last hour:", len(schds.Data))
//
type ListRecipientTransferSchedules struct {
	RecipientID string
	List
}

func (req *ListRecipientTransferSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients/" + req.RecipientID + "/schedules",
		ContentType: "application/json",
	}
}
