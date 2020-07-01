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
//	dispute, err := client.Dispute().ListClosed(ctx, &ListClosedDisputeParams{
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().List(ctx, &ListDisputesParams{
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().ListOpen(ctx, &ListOpenDisputeParams{
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().ListPending(ctx, &ListPendingDisputeParams{
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().Retrieve(ctx, &RetrieveDisputeParams{
//		DisputeID: "dspt_9c4h4hr1eo"
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().Update(ctx, &UpdateDisputeParams{
//		DisputeID: "dspt_9c4h4hr1eo"
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().Close(ctx, &CloseDisputeParams{
//		DisputeID: "dspt_9c4h4hr1eo"
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().Accept(ctx, &AcceptDisputeParams{
//		DisputeID: "dspt_9c4h4hr1eo"
//	})
//	if err != nil {
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
//	dispute, err := client.Dispute().Create(ctx, &CreateDisputeParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
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

