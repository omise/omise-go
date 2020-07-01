package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Charge represents Charge object.
// See https://www.omise.co/charges-api for more information.
type Charge struct {
	Base
	Amount int64 `json:"amount"`
	AuthorizeURI *string `json:"authorize_uri"`
	Authorized bool `json:"authorized"`
	Branch string `json:"branch"`
	Capturable bool `json:"capturable"`
	Capture bool `json:"capture"`
	Card *Card `json:"card"`
	Currency string `json:"currency"`
	Customer string `json:"customer"`
	Description *string `json:"description"`
	Device string `json:"device"`
	Disputable bool `json:"disputable"`
	Dispute *Dispute `json:"dispute"`
	Expired bool `json:"expired"`
	ExpiredAt *time.Time `json:"expired_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	FailureCode *ChargeFailureCode `json:"failure_code"`
	FailureMessage *string `json:"failure_message"`
	Fee int64 `json:"fee"`
	FeeVat int64 `json:"fee_vat"`
	FundingAmount int64 `json:"funding_amount"`
	FundingCurrency string `json:"funding_currency"`
	Interest int64 `json:"interest"`
	InterestVat int64 `json:"interest_vat"`
	IP *string `json:"ip"`
	Link string `json:"link"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Net int64 `json:"net"`
	Paid bool `json:"paid"`
	PaidAt time.Time `json:"paid_at"`
	PlatformFee *PlatformFee `json:"platform_fee"`
	Refundable bool `json:"refundable"`
	RefundedAmount int64 `json:"refunded_amount"`
	Refunds *RefundList `json:"refunds"`
	ReturnURI *string `json:"return_uri"`
	Reversed bool `json:"reversed"`
	ReversedAt *time.Time `json:"reversed_at"`
	Reversible bool `json:"reversible"`
	Schedule string `json:"schedule"`
	Source *Source `json:"source"`
	Status ChargeStatus `json:"status"`
	Terminal string `json:"terminal"`
	Transaction string `json:"transaction"`
	Voided bool `json:"voided"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}

// ChargeService represents resource service.
type ChargeService struct {
	*Client
}

// Charge defines resource service.
func (c *Client) Charge() *ChargeService {
	return &ChargeService{c}
}

// List lists charges
//
// Example:
//
//	charge, err := client.Charge().List(ctx, &ListChargesParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) List(ctx context.Context, params *ListChargesParams) (*ChargeList, error) {
	result := &ChargeList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListChargesParams params object.
type ListChargesParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListChargesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges",
		ContentType: "application/json",
	}
}

// Create creates charge
//
// Example:
//
//	charge, err := client.Charge().Create(ctx, &CreateChargeParams{
//		Amount: 10000
//		Currency: "thb"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) Create(ctx context.Context, params *CreateChargeParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateChargeParams params object.
type CreateChargeParams struct {
	Amount int64 `json:"amount"`
	Capture *bool `json:"capture,omitempty"`
	Card string `json:"card,omitempty"`
	Currency string `json:"currency"`
	Customer string `json:"customer,omitempty"`
	Description string `json:"description,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	IP string `json:"ip,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	PlatformFee *PlatformFeeParams `json:"platform_fee,omitempty"`
	Reference string `json:"reference,omitempty"`
	ReturnURI string `json:"return_uri,omitempty"`
	Source string `json:"source,omitempty"`
}

// PlatformFeeParams params object.
type PlatformFeeParams struct {
	Amount int64 `json:"amount,omitempty"`
	Fixed int64 `json:"fixed,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}

// Describe describes structure of request
func (req *CreateChargeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges",
		ContentType: "application/json",
	}
}

// ListEvents lists charge events
//
// Example:
//
//	event, err := client.Charge().ListEvents(ctx, &ListChargeEventsParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("event: %#v\n", event)
//
func (s *ChargeService) ListEvents(ctx context.Context, params *ListChargeEventsParams) (*EventList, error) {
	result := &EventList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListChargeEventsParams params object.
type ListChargeEventsParams struct {
	ChargeID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListChargeEventsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID + "/events",
		ContentType: "application/json",
	}
}

// ListSchedules lists charge schedules
//
// Example:
//
//	charge_schedule, err := client.Charge().ListSchedules(ctx, &ListChargeSchedulesParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge_schedule: %#v\n", charge_schedule)
//
func (s *ChargeService) ListSchedules(ctx context.Context, params *ListChargeSchedulesParams) (*ChargeScheduleList, error) {
	result := &ChargeScheduleList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListChargeSchedulesParams params object.
type ListChargeSchedulesParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListChargeSchedulesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/schedules",
		ContentType: "application/json",
	}
}

// Retrieve retrieves charge
//
// Example:
//
//	charge, err := client.Charge().Retrieve(ctx, &RetrieveChargeParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) Retrieve(ctx context.Context, params *RetrieveChargeParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveChargeParams params object.
type RetrieveChargeParams struct {
	ChargeID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveChargeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
	}
}

// Update updates charge
//
// Example:
//
//	charge, err := client.Charge().Update(ctx, &UpdateChargeParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) Update(ctx context.Context, params *UpdateChargeParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// UpdateChargeParams params object.
type UpdateChargeParams struct {
	ChargeID string `json:"-"`
	Description string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe describes structure of request
func (req *UpdateChargeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
	}
}

// Capture captures charge
//
// Example:
//
//	charge, err := client.Charge().Capture(ctx, &CaptureChargeParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) Capture(ctx context.Context, params *CaptureChargeParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CaptureChargeParams params object.
type CaptureChargeParams struct {
	ChargeID string `json:"-"`
}

// Describe describes structure of request
func (req *CaptureChargeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/capture",
		ContentType: "application/json",
	}
}

// Expire expires charge
//
// Example:
//
//	charge, err := client.Charge().Expire(ctx, &ExpireChargeParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) Expire(ctx context.Context, params *ExpireChargeParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ExpireChargeParams params object.
type ExpireChargeParams struct {
	ChargeID string `json:"-"`
}

// Describe describes structure of request
func (req *ExpireChargeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/expire",
		ContentType: "application/json",
	}
}

// MarkAsFailed marks charge as failed
//
// Example:
//
//	charge, err := client.Charge().MarkAsFailed(ctx, &MarkChargeAsFailedParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) MarkAsFailed(ctx context.Context, params *MarkChargeAsFailedParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// MarkChargeAsFailedParams params object.
type MarkChargeAsFailedParams struct {
	ChargeID string `json:"-"`
}

// Describe describes structure of request
func (req *MarkChargeAsFailedParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/mark_as_failed",
		ContentType: "application/json",
	}
}

// MarkAsPaid marks charge as paid
//
// Example:
//
//	charge, err := client.Charge().MarkAsPaid(ctx, &MarkChargeAsPaidParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) MarkAsPaid(ctx context.Context, params *MarkChargeAsPaidParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// MarkChargeAsPaidParams params object.
type MarkChargeAsPaidParams struct {
	ChargeID string `json:"-"`
}

// Describe describes structure of request
func (req *MarkChargeAsPaidParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/mark_as_paid",
		ContentType: "application/json",
	}
}

// Reverse reverses charge
//
// Example:
//
//	charge, err := client.Charge().Reverse(ctx, &ReverseChargeParams{
//		ChargeID: "chrg_8b3g63gq2f"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *ChargeService) Reverse(ctx context.Context, params *ReverseChargeParams) (*Charge, error) {
	result := &Charge{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ReverseChargeParams params object.
type ReverseChargeParams struct {
	ChargeID string `json:"-"`
}

// Describe describes structure of request
func (req *ReverseChargeParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/reverse",
		ContentType: "application/json",
	}
}

