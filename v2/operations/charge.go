package operations

import (
	"time"
)

// Example:
//
//	charge, update := &omise.Charge{}, &SearchCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchCharge struct {
	Base
}

func (req *SearchCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/charges/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &ListCharges{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListCharges struct {
	Base
}

func (req *ListCharges) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/charges",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &CreateCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateCharge struct {
	Base
	Amount int `json:"amount"`
	Capture bool `json:"capture"`
	Card string `json:"card"`
	Currency string `json:"currency"`
	Customer string `json:"customer"`
	Description string `json:"description"`
	ExpiresAt time.Time `json:"expires_at"`
	IP string `json:"ip"`
	Metadata map[string]interface{} `json:"metadata"`
	PlatformFee *PlatformFee `json:"platform_fee"`
	Reference string `json:"reference"`
	ReturnURI string `json:"return_uri"`
	Source string `json:"source"`
}

func (req *CreateCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Event{}, &ListChargesEvents{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListChargesEvents struct {
	Base
	ChargeID string `json:"-"`
}

func (req *ListChargesEvents) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/charges/" + req.ChargeID + "/events",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &RetrieveCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveCharge struct {
	Base
	ChargeID string `json:"-"`
}

func (req *RetrieveCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &UpdateCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateCharge struct {
	Base
	ChargeID string `json:"-"`
	Description string `json:"description"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (req *UpdateCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      PATCH,
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &CaptureCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CaptureCharge struct {
	Base
	ChargeID string `json:"-"`
}

func (req *CaptureCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges/" + req.ChargeID + "/capture",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &ExpireCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ExpireCharge struct {
	Base
	ChargeID string `json:"-"`
}

func (req *ExpireCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges/" + req.ChargeID + "/expire",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &MarkChargeAsFailed{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type MarkChargeAsFailed struct {
	Base
	ChargeID string `json:"-"`
}

func (req *MarkChargeAsFailed) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges/" + req.ChargeID + "/mark_as_failed",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &MarkChargeAsPaid{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type MarkChargeAsPaid struct {
	Base
	ChargeID string `json:"-"`
}

func (req *MarkChargeAsPaid) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges/" + req.ChargeID + "/mark_as_paid",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &ReverseCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ReverseCharge struct {
	Base
	ChargeID string `json:"-"`
}

func (req *ReverseCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges/" + req.ChargeID + "/reverse",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.ChargeSchedule{}, &SearchChargeSchedule{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchChargeSchedule struct {
	Base
}

func (req *SearchChargeSchedule) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/charges/schedules/" + "/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &CreateChargeDisputeDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateChargeDisputeDispute struct {
	Base
	ChargeID string `json:"-"`
}

func (req *CreateChargeDisputeDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges/" + req.ChargeID + "/disputes",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Refund{}, &ListChargesRefund{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListChargesRefund struct {
	Base
	ChargeID string `json:"-"`
}

func (req *ListChargesRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/charges/" + req.ChargeID + "/refunds",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Refund{}, &CreateChargeRefund{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateChargeRefund struct {
	Base
	ChargeID string `json:"-"`
	Amount int `json:"amount"`
	Metadata map[string]interface{} `json:"metadata"`
	Void bool `json:"void"`
}

func (req *CreateChargeRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/charges/" + req.ChargeID + "/refunds",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Refund{}, &RetrieveChargeRefund{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveChargeRefund struct {
	Base
	ChargeID string `json:"-"`
	RefundID string `json:"-"`
}

func (req *RetrieveChargeRefund) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/charges/" + req.ChargeID + "/refunds/" + req.RefundID,
		ContentType: "application/json",
	}
}

