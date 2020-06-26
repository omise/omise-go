package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Refund represents Refund object.
// See https://www.omise.co/refunds-api for more information.
type Refund struct {
	Base
	Amount int64 `json:"amount"`
	Charge string `json:"charge"`
	Currency string `json:"currency"`
	FundingAmount int64 `json:"funding_amount"`
	FundingCurrency string `json:"funding_currency"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Status RefundStatus `json:"status"`
	Transaction string `json:"transaction"`
	Voided bool `json:"voided"`
}

// RefundService represents resource service.
type RefundService struct {
	*Client
}

// Refund defines resource service.
func (c *Client) Refund() *RefundService {
	return &RefundService{c}
}

// ListRefunds lists all refunds
//
// Example:
//
//	refund, listRefunds := &omise.RefundList{}, &ListAllRefunds{
//	}
//	if e := client.Do(refund, listRefunds); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
func (s *RefundService) ListRefunds(ctx context.Context, params *ListAllRefundsParams) (*RefundList, error) {
	result := &RefundList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListAllRefundsParams params object.
type ListAllRefundsParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListAllRefundsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/refunds",
		ContentType: "application/json",
	}
}

// List lists refunds
//
// Example:
//
//	refund, list := &omise.RefundList{}, &ListRefunds{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(refund, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
func (s *RefundService) List(ctx context.Context, params *ListRefundsParams) (*RefundList, error) {
	result := &RefundList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListRefundsParams params object.
type ListRefundsParams struct {
	ChargeID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListRefundsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID + "/refunds",
		ContentType: "application/json",
	}
}

// Create creates refund
//
// Example:
//
//	refund, create := &omise.Refund{}, &CreateRefund{
//		ChargeID: "chrg_8b3g63gq2f"
//		Amount: 10000
//	}
//	if e := client.Do(refund, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
func (s *RefundService) Create(ctx context.Context, params *CreateRefundParams) (*Refund, error) {
	result := &Refund{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateRefundParams params object.
type CreateRefundParams struct {
	ChargeID string `json:"-"`
	Amount int64 `json:"amount"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Void bool `json:"void,omitempty"`
}

// Describe describes structure of request
func (req *CreateRefundParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/refunds",
		ContentType: "application/json",
	}
}

// Retrieve retrieves refund
//
// Example:
//
//	refund, retrieve := &omise.Refund{}, &RetrieveRefund{
//		ChargeID: "chrg_8b3g63gq2f"
//		RefundID: "refun_3gqd1e6jt9"
//	}
//	if e := client.Do(refund, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("refund: %#v\n", refund)
//
func (s *RefundService) Retrieve(ctx context.Context, params *RetrieveRefundParams) (*Refund, error) {
	result := &Refund{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveRefundParams params object.
type RetrieveRefundParams struct {
	ChargeID string `json:"-"`
	RefundID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveRefundParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID + "/refunds/" + req.RefundID,
		ContentType: "application/json",
	}
}

