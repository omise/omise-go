package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Transfer represents Transfer object.
// See https://www.omise.co/transfers-api for more information.
type Transfer struct {
	Base
	Amount int64 `json:"amount"`
	BankAccount *BankAccount `json:"bank_account"`
	Currency string `json:"currency"`
	FailFast bool `json:"fail_fast"`
	FailureCode *string `json:"failure_code"`
	FailureMessage *string `json:"failure_message"`
	Fee int64 `json:"fee"`
	FeeVat int64 `json:"fee_vat"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Net int64 `json:"net"`
	Paid bool `json:"paid"`
	PaidAt *time.Time `json:"paid_at"`
	Recipient string `json:"recipient"`
	Schedule string `json:"schedule"`
	Sendable bool `json:"sendable"`
	Sent bool `json:"sent"`
	SentAt *time.Time `json:"sent_at"`
	TotalFee int64 `json:"total_fee"`
	Transactions []Transaction `json:"transactions"`
}

// TransferService represents resource service.
type TransferService struct {
	*Client
}

// Transfer defines resource service.
func (c *Client) Transfer() *TransferService {
	return &TransferService{c}
}

// Destroy destroys transfer
//
// Example:
//
//	transfer, destroy := &omise.Transfer{}, &DestroyTransfer{
//		TransferID: "trsf_5is3gq69cm"
//	}
//	if e := client.Do(transfer, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
func (s *TransferService) Destroy(ctx context.Context, params *DestroyTransferParams) (*Transfer, error) {
	result := &Transfer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DestroyTransferParams params object.
type DestroyTransferParams struct {
	TransferID string `json:"-"`
}

// Describe describes structure of request
func (req *DestroyTransferParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

// Retrieve retrieves transfer
//
// Example:
//
//	transfer, retrieve := &omise.Transfer{}, &RetrieveTransfer{
//		TransferID: "trsf_5is3gq69cm"
//	}
//	if e := client.Do(transfer, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
func (s *TransferService) Retrieve(ctx context.Context, params *RetrieveTransferParams) (*Transfer, error) {
	result := &Transfer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveTransferParams params object.
type RetrieveTransferParams struct {
	TransferID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveTransferParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

// Update updates transfer
//
// Example:
//
//	transfer, update := &omise.Transfer{}, &UpdateTransfer{
//		TransferID: "trsf_5is3gq69cm"
//	}
//	if e := client.Do(transfer, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
func (s *TransferService) Update(ctx context.Context, params *UpdateTransferParams) (*Transfer, error) {
	result := &Transfer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// UpdateTransferParams params object.
type UpdateTransferParams struct {
	TransferID string `json:"-"`
	Amount int64 `json:"amount,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe describes structure of request
func (req *UpdateTransferParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

// List lists transfers
//
// Example:
//
//	transfer, list := &omise.TransferList{}, &ListTransfers{
//	}
//	if e := client.Do(transfer, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
func (s *TransferService) List(ctx context.Context, params *ListTransfersParams) (*TransferList, error) {
	result := &TransferList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListTransfersParams params object.
type ListTransfersParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListTransfersParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transfers",
		ContentType: "application/json",
	}
}

// Create creates transfer
//
// Example:
//
//	transfer, create := &omise.Transfer{}, &CreateTransfer{
//	}
//	if e := client.Do(transfer, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
func (s *TransferService) Create(ctx context.Context, params *CreateTransferParams) (*Transfer, error) {
	result := &Transfer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateTransferParams params object.
type CreateTransferParams struct {
	Amount int64 `json:"amount,omitempty"`
	FailFast bool `json:"fail_fast,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Recipient string `json:"recipient,omitempty"`
}

// Describe describes structure of request
func (req *CreateTransferParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/transfers",
		ContentType: "application/json",
	}
}

// ListSchedules lists transfer schedules
//
// Example:
//
//	transfer_schedule, listSchedules := &omise.TransferScheduleList{}, &ListTransferSchedules{
//	}
//	if e := client.Do(transfer_schedule, listSchedules); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer_schedule: %#v\n", transfer_schedule)
//
func (s *TransferService) ListSchedules(ctx context.Context, params *ListTransferSchedulesParams) (*TransferScheduleList, error) {
	result := &TransferScheduleList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListTransferSchedulesParams params object.
type ListTransferSchedulesParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListTransferSchedulesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transfers/schedules",
		ContentType: "application/json",
	}
}

// MarkAsPaid marks transfer as paid
//
// Example:
//
//	transfer, markAsPaid := &omise.Transfer{}, &MarkTransferAsPaid{
//		TransferID: "trsf_5is3gq69cm"
//	}
//	if e := client.Do(transfer, markAsPaid); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
func (s *TransferService) MarkAsPaid(ctx context.Context, params *MarkTransferAsPaidParams) (*Transfer, error) {
	result := &Transfer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// MarkTransferAsPaidParams params object.
type MarkTransferAsPaidParams struct {
	TransferID string `json:"-"`
}

// Describe describes structure of request
func (req *MarkTransferAsPaidParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/transfers/" + req.TransferID + "/mark_as_paid",
		ContentType: "application/json",
	}
}

// MarkAsSent marks transfer as sent
//
// Example:
//
//	transfer, markAsSent := &omise.Transfer{}, &MarkTransferAsSent{
//		TransferID: "trsf_5is3gq69cm"
//	}
//	if e := client.Do(transfer, markAsSent); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
func (s *TransferService) MarkAsSent(ctx context.Context, params *MarkTransferAsSentParams) (*Transfer, error) {
	result := &Transfer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// MarkTransferAsSentParams params object.
type MarkTransferAsSentParams struct {
	TransferID string `json:"-"`
}

// Describe describes structure of request
func (req *MarkTransferAsSentParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/transfers/" + req.TransferID + "/mark_as_sent",
		ContentType: "application/json",
	}
}

