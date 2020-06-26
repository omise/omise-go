package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Dispute represents Dispute object.
// See https://www.omise.co/disputes-api for more information.
type Dispute struct {
	Base
	AdminMessage *string `json:"admin_message"`
	Amount int64 `json:"amount"`
	Charge string `json:"charge"`
	ClosedAt *time.Time `json:"closed_at"`
	Currency string `json:"currency"`
	Documents *DocumentList `json:"documents"`
	FundingAmount int64 `json:"funding_amount"`
	FundingCurrency string `json:"funding_currency"`
	Location string `json:"location"`
	Message string `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
	ReasonCode DisputeReasonCode `json:"reason_code"`
	ReasonMessage string `json:"reason_message"`
	Status DisputeStatus `json:"status"`
	Transactions []Transaction `json:"transactions"`
}

// DisputeService represents resource service.
type DisputeService struct {
	*Client
}

// Dispute defines resource service.
func (c *Client) Dispute() *DisputeService {
	return &DisputeService{c}
}

// ListClosed closeds dispute
//
// Example:
//
//	dispute, closed := &omise.DisputeList{}, &ListClosedDispute{
//	}
//	if e := client.Do(dispute, closed); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) ListClosed(ctx context.Context, params *ListClosedDisputeParams) (*DisputeList, error) {
	result := &DisputeList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListClosedDisputeParams params object.
type ListClosedDisputeParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListClosedDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/closed",
		ContentType: "application/json",
	}
}

// List lists disputes
//
// Example:
//
//	dispute, list := &omise.DisputeList{}, &ListDisputes{
//	}
//	if e := client.Do(dispute, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) List(ctx context.Context, params *ListDisputesParams) (*DisputeList, error) {
	result := &DisputeList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListDisputesParams params object.
type ListDisputesParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListDisputesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes",
		ContentType: "application/json",
	}
}

// ListOpen opens dispute
//
// Example:
//
//	dispute, open := &omise.DisputeList{}, &ListOpenDispute{
//	}
//	if e := client.Do(dispute, open); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) ListOpen(ctx context.Context, params *ListOpenDisputeParams) (*DisputeList, error) {
	result := &DisputeList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListOpenDisputeParams params object.
type ListOpenDisputeParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListOpenDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/open",
		ContentType: "application/json",
	}
}

// ListPending pendings dispute
//
// Example:
//
//	dispute, pending := &omise.DisputeList{}, &ListPendingDispute{
//	}
//	if e := client.Do(dispute, pending); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) ListPending(ctx context.Context, params *ListPendingDisputeParams) (*DisputeList, error) {
	result := &DisputeList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListPendingDisputeParams params object.
type ListPendingDisputeParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListPendingDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/pending",
		ContentType: "application/json",
	}
}

// Retrieve retrieves dispute
//
// Example:
//
//	dispute, retrieve := &omise.Dispute{}, &RetrieveDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) Retrieve(ctx context.Context, params *RetrieveDisputeParams) (*Dispute, error) {
	result := &Dispute{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveDisputeParams params object.
type RetrieveDisputeParams struct {
	DisputeID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID,
		ContentType: "application/json",
	}
}

// Update updates dispute
//
// Example:
//
//	dispute, update := &omise.Dispute{}, &UpdateDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) Update(ctx context.Context, params *UpdateDisputeParams) (*Dispute, error) {
	result := &Dispute{}
	err := s.Do(ctx, result, params)

	return result, err
}

// UpdateDisputeParams params object.
type UpdateDisputeParams struct {
	DisputeID string `json:"-"`
	Message string `json:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe describes structure of request
func (req *UpdateDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/disputes/" + req.DisputeID,
		ContentType: "application/json",
	}
}

// Close closes dispute
//
// Example:
//
//	dispute, close := &omise.Dispute{}, &CloseDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, close); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) Close(ctx context.Context, params *CloseDisputeParams) (*Dispute, error) {
	result := &Dispute{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CloseDisputeParams params object.
type CloseDisputeParams struct {
	DisputeID string `json:"-"`
	Status DisputeStatus `json:"status,omitempty"`
}

// Describe describes structure of request
func (req *CloseDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/disputes/" + req.DisputeID + "/close",
		ContentType: "application/json",
	}
}

// Accept accepts dispute
//
// Example:
//
//	dispute, accept := &omise.Dispute{}, &AcceptDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, accept); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) Accept(ctx context.Context, params *AcceptDisputeParams) (*Dispute, error) {
	result := &Dispute{}
	err := s.Do(ctx, result, params)

	return result, err
}

// AcceptDisputeParams params object.
type AcceptDisputeParams struct {
	DisputeID string `json:"-"`
}

// Describe describes structure of request
func (req *AcceptDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/disputes/" + req.DisputeID + "/accept",
		ContentType: "application/json",
	}
}

// Create creates dispute
//
// Example:
//
//	dispute, createDispute := &omise.Dispute{}, &CreateDispute{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(dispute, createDispute); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
func (s *DisputeService) Create(ctx context.Context, params *CreateDisputeParams) (*Dispute, error) {
	result := &Dispute{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateDisputeParams params object.
type CreateDisputeParams struct {
	ChargeID string `json:"-"`
}

// Describe describes structure of request
func (req *CreateDisputeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/disputes",
		ContentType: "application/json",
	}
}

