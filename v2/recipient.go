package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Recipient represents Recipient object.
// See https://www.omise.co/recipients-api for more information.
type Recipient struct {
	Base
	ActivatedAt time.Time `json:"activated_at"`
	Active bool `json:"active"`
	BankAccount *BankAccount `json:"bank_account"`
	Default bool `json:"default"`
	Description *string `json:"description"`
	Email string `json:"email"`
	FailureCode *RecipientFailureCode `json:"failure_code"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Name string `json:"name"`
	Schedule string `json:"schedule"`
	TaxID *string `json:"tax_id"`
	Type RecipientType `json:"type"`
	Verified bool `json:"verified"`
	VerifiedAt time.Time `json:"verified_at"`
}

// RecipientService represents resource service.
type RecipientService struct {
	*Client
}

// Recipient defines resource service.
func (c *Client) Recipient() *RecipientService {
	return &RecipientService{c}
}

// Destroy destroys recipient
//
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
func (s *RecipientService) Destroy(ctx context.Context, params *DestroyRecipientParams) (*Recipient, error) {
	result := &Recipient{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DestroyRecipientParams params object.
type DestroyRecipientParams struct {
	RecipientID string `json:"-"`
}

// Describe describes structure of request
func (req *DestroyRecipientParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Retrieve retrieves recipient
//
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
func (s *RecipientService) Retrieve(ctx context.Context, params *RetrieveRecipientParams) (*Recipient, error) {
	result := &Recipient{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveRecipientParams params object.
type RetrieveRecipientParams struct {
	RecipientID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveRecipientParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Update updates recipient
//
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
func (s *RecipientService) Update(ctx context.Context, params *UpdateRecipientParams) (*Recipient, error) {
	result := &Recipient{}
	err := s.Do(ctx, result, params)

	return result, err
}

// UpdateRecipientParams params object.
type UpdateRecipientParams struct {
	RecipientID string `json:"-"`
	BankAccount *BankAccountParams `json:"bank_account,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name,omitempty"`
	TaxID string `json:"tax_id,omitempty"`
	Type RecipientType `json:"type,omitempty"`
}

// BankAccountParams params object.
type BankAccountParams struct {
	BankCode string `json:"bank_code,omitempty"`
	BranchCode string `json:"branch_code,omitempty"`
	Brand string `json:"brand,omitempty"`
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	Type BankAccountType `json:"type,omitempty"`
}

// Describe describes structure of request
func (req *UpdateRecipientParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// List lists recipients
//
// Example:
//
//	recipient, list := &omise.RecipientList{}, &ListRecipients{
//	}
//	if e := client.Do(recipient, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("recipient: %#v\n", recipient)
//
func (s *RecipientService) List(ctx context.Context, params *ListRecipientsParams) (*RecipientList, error) {
	result := &RecipientList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListRecipientsParams params object.
type ListRecipientsParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListRecipientsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients",
		ContentType: "application/json",
	}
}

// Create creates recipient
//
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
func (s *RecipientService) Create(ctx context.Context, params *CreateRecipientParams) (*Recipient, error) {
	result := &Recipient{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateRecipientParams params object.
type CreateRecipientParams struct {
	BankAccount *BankAccountParams `json:"bank_account,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name,omitempty"`
	TaxID string `json:"tax_id,omitempty"`
	Type RecipientType `json:"type,omitempty"`
}

// Describe describes structure of request
func (req *CreateRecipientParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/recipients",
		ContentType: "application/json",
	}
}

// ListSchedules lists recipient schedules
//
// Example:
//
//	schedule, listSchedules := &omise.ScheduleList{}, &ListRecipientSchedules{
//		RecipientID: "recip_3gqd8b4h1e"
//	}
//	if e := client.Do(schedule, listSchedules); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
func (s *RecipientService) ListSchedules(ctx context.Context, params *ListRecipientSchedulesParams) (*ScheduleList, error) {
	result := &ScheduleList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListRecipientSchedulesParams params object.
type ListRecipientSchedulesParams struct {
	RecipientID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListRecipientSchedulesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/recipients/" + req.RecipientID + "/schedules",
		ContentType: "application/json",
	}
}

// Verify verifys recipient
//
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
func (s *RecipientService) Verify(ctx context.Context, params *VerifyRecipientParams) (*Recipient, error) {
	result := &Recipient{}
	err := s.Do(ctx, result, params)

	return result, err
}

// VerifyRecipientParams params object.
type VerifyRecipientParams struct {
	RecipientID string `json:"-"`
}

// Describe describes structure of request
func (req *VerifyRecipientParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/recipients/" + req.RecipientID + "/verify",
		ContentType: "application/json",
	}
}

