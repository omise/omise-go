package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Dispute{}, &SearchDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchDispute struct {
}

func (req *SearchDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &ClosedDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ClosedDispute struct {
}

func (req *ClosedDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/closed",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &ListDisputes{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListDisputes struct {
}

func (req *ListDisputes) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &OpenDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type OpenDispute struct {
}

func (req *OpenDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/open",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &PendingDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type PendingDispute struct {
}

func (req *PendingDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/pending",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &RetrieveDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveDispute struct {
	DisputeID string `json:"-"`
}

func (req *RetrieveDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &UpdateDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateDispute struct {
	DisputeID string `json:"-"`
	Message string `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (req *UpdateDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/disputes/" + req.DisputeID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &CloseDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CloseDispute struct {
	DisputeID string `json:"-"`
	Status *DisputeStatus `json:"status"`
}

func (req *CloseDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/disputes/" + req.DisputeID + "/close",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Dispute{}, &AcceptDispute{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type AcceptDispute struct {
	DisputeID string `json:"-"`
}

func (req *AcceptDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/disputes/" + req.DisputeID + "/accept",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Document{}, &DestroyDisputeDocument{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyDisputeDocument struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

func (req *DestroyDisputeDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Document{}, &RetrieveDisputeDocument{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveDisputeDocument struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

func (req *RetrieveDisputeDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Document{}, &ListDisputesDocument{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListDisputesDocument struct {
	DisputeID string `json:"-"`
}

func (req *ListDisputesDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID + "/documents",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Document{}, &CreateDisputeDocument{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateDisputeDocument struct {
	DisputeID string `json:"-"`
	File string `json:"file"`
}

func (req *CreateDisputeDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/disputes/" + req.DisputeID + "/documents",
		ContentType: "application/json",
	}
}

