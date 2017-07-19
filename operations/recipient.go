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
	Name        string
	Email       string
	Description string
	Type        omise.RecipientType
	TaxID       string             `query:"tax_id"`
	BankAccount *omise.BankAccount `query:"bank_account"`
}

func (req *CreateRecipient) MarshalJSON() ([]byte, error) {
	type bankAccount struct {
		Brand  string `json:"brand,omitempty"`
		Number string `json:"number"`
		Name   string `json:"name"`

		// for Omise Japan
		BankCode    string                `json:"bank_code,omitempty"`
		BranchCode  string                `json:"branch_code,omitempty"`
		AccountType omise.BankAccountType `json:"account_type,omitempty"`
	}

	type param struct {
		Name        string              `json:"name"`
		Email       string              `json:"email,omitempty"`
		Description string              `json:"description,omitempty"`
		Type        omise.RecipientType `json:"type"`
		TaxID       string              `json:"tax_id,omitempty"`
		BankAccount *bankAccount        `json:"bank_account,omitempty"`
	}

	p := param{
		Name:        req.Name,
		Email:       req.Email,
		Description: req.Description,
		Type:        req.Type,
		TaxID:       req.TaxID,
	}
	if req.BankAccount != nil {
		p.BankAccount = &bankAccount{
			Brand:       req.BankAccount.Brand,
			Number:      req.BankAccount.Number,
			Name:        req.BankAccount.Name,
			BankCode:    req.BankAccount.BankCode,
			BranchCode:  req.BankAccount.BranchCode,
			AccountType: req.BankAccount.AccountType,
		}
	}

	return json.Marshal(p)
}

func (req *CreateRecipient) Op() *internal.Op {
	return &internal.Op{
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

func (req *UpdateRecipient) MarshalJSON() ([]byte, error) {
	type bankAccount struct {
		Brand  string `json:"brand,omitempty"`
		Number string `json:"number"`
		Name   string `json:"name"`

		// for Omise Japan
		BankCode    string                `json:"bank_code,omitempty"`
		BranchCode  string                `json:"branch_code,omitempty"`
		AccountType omise.BankAccountType `json:"account_type,omitempty"`
	}

	type param struct {
		Name        string              `json:"name,omitempty"`
		Email       string              `json:"email,omitempty"`
		Description string              `json:"description,omitempty"`
		Type        omise.RecipientType `json:"type,omitempty"`
		TaxID       string              `json:"tax_id,omitempty"`
		BankAccount *bankAccount        `json:"bank_account,omitempty"`
	}

	p := param{
		Name:        req.Name,
		Email:       req.Email,
		Description: req.Description,
		Type:        req.Type,
		TaxID:       req.TaxID,
	}
	if req.BankAccount != nil {
		p.BankAccount = &bankAccount{
			Brand:       req.BankAccount.Brand,
			Number:      req.BankAccount.Number,
			Name:        req.BankAccount.Name,
			BankCode:    req.BankAccount.BankCode,
			BranchCode:  req.BankAccount.BranchCode,
			AccountType: req.BankAccount.AccountType,
		}
	}

	return json.Marshal(p)
}

func (req *UpdateRecipient) Op() *internal.Op {
	return &internal.Op{
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
	RecipientID string `query:"-"`
}

func (req *DestroyRecipient) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "DELETE",
		Path:     "/recipients/" + req.RecipientID,
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

func (req *ListRecipientTransferSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(req.List)
}

func (req *ListRecipientTransferSchedules) Op() *internal.Op {
	return &internal.Op{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients/" + req.RecipientID + "/schedules",
		ContentType: "application/json",
	}
}
