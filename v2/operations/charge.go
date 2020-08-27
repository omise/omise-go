package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"time"
)

// Example:
//
//	charge, list := &omise.Charge{}, &ListCharges{
//	}
//	if e := client.Do(charge, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type ListCharges struct {
	List
}

// Describe
func (req *ListCharges) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, create := &omise.Charge{}, &CreateCharge{
//		Amount: 10000
//		Currency: "thb"
//	}
//	if e := client.Do(charge, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type CreateCharge struct {
	Amount int64 `json:"amount"`
	Capture *bool `json:"capture,omitempty"`
	Card string `json:"card,omitempty"`
	Currency string `json:"currency"`
	Customer string `json:"customer,omitempty"`
	Description string `json:"description,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	IP string `json:"ip,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	PlatformFee *PlatformFee `json:"platform_fee,omitempty"`
	Reference string `json:"reference,omitempty"`
	ReturnURI string `json:"return_uri,omitempty"`
	Source string `json:"source,omitempty"`
}

// PlatformFee
type PlatformFee struct {
	Amount int64 `json:"amount,omitempty"`
	Fixed int64 `json:"fixed,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}

// Describe
func (req *CreateCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges",
		ContentType: "application/json",
	}
}

// Example:
//
//	event, listEvents := &omise.Event{}, &ListChargeEvents{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(event, listEvents); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("event: %#v\n", event)
//
type ListChargeEvents struct {
	ChargeID string `json:"-"`
	List
}

// Describe
func (req *ListChargeEvents) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID + "/events",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge_schedule, listSchedules := &omise.ChargeSchedule{}, &ListChargeSchedules{
//	}
//	if e := client.Do(charge_schedule, listSchedules); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge_schedule: %#v\n", charge_schedule)
//
type ListChargeSchedules struct {
	List
}

// Describe
func (req *ListChargeSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, retrieve := &omise.Charge{}, &RetrieveCharge{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(charge, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type RetrieveCharge struct {
	ChargeID string `json:"-"`
}

// Describe
func (req *RetrieveCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Charge{}, &UpdateCharge{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type UpdateCharge struct {
	ChargeID string `json:"-"`
	Description string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe
func (req *UpdateCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, capture := &omise.Charge{}, &CaptureCharge{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(charge, capture); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type CaptureCharge struct {
	ChargeID string `json:"-"`
}

// Describe
func (req *CaptureCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/capture",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, expire := &omise.Charge{}, &ExpireCharge{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(charge, expire); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type ExpireCharge struct {
	ChargeID string `json:"-"`
}

// Describe
func (req *ExpireCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/expire",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, markAsFailed := &omise.Charge{}, &MarkChargeAsFailed{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(charge, markAsFailed); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type MarkChargeAsFailed struct {
	ChargeID string `json:"-"`
}

// Describe
func (req *MarkChargeAsFailed) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/mark_as_failed",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, markAsPaid := &omise.Charge{}, &MarkChargeAsPaid{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(charge, markAsPaid); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type MarkChargeAsPaid struct {
	ChargeID string `json:"-"`
}

// Describe
func (req *MarkChargeAsPaid) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/mark_as_paid",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, reverse := &omise.Charge{}, &ReverseCharge{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(charge, reverse); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type ReverseCharge struct {
	ChargeID string `json:"-"`
}

// Describe
func (req *ReverseCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/reverse",
		ContentType: "application/json",
	}
}

