package operations

import (
	"github.com/omise/omise-go/v2/internal"
	"github.com/omise/omise-go/v2"
)

// Example:
//
//	dispute, closed := &omise.Dispute{}, &ClosedDispute{
//	}
//	if e := client.Do(dispute, closed); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type ClosedDispute struct {
}

// Describe
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
//	dispute, list := &omise.Dispute{}, &ListDisputes{
//	}
//	if e := client.Do(dispute, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type ListDisputes struct {
	List
}

// Describe
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
//	dispute, open := &omise.Dispute{}, &OpenDispute{
//	}
//	if e := client.Do(dispute, open); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type OpenDispute struct {
}

// Describe
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
//	dispute, pending := &omise.Dispute{}, &PendingDispute{
//	}
//	if e := client.Do(dispute, pending); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type PendingDispute struct {
}

// Describe
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
//	dispute, retrieve := &omise.Dispute{}, &RetrieveDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type RetrieveDispute struct {
	DisputeID string `json:"-"`
}

// Describe
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
//	dispute, update := &omise.Dispute{}, &UpdateDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type UpdateDispute struct {
	DisputeID string `json:"-"`
	Message string `json:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe
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
//	dispute, close := &omise.Dispute{}, &CloseDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, close); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type CloseDispute struct {
	DisputeID string `json:"-"`
	Status omise.DisputeStatus `json:"status,omitempty"`
}

// Describe
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
//	dispute, accept := &omise.Dispute{}, &AcceptDispute{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(dispute, accept); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type AcceptDispute struct {
	DisputeID string `json:"-"`
}

// Describe
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
//	dispute, createDispute := &omise.Dispute{}, &CreateDispute{
//		ChargeID: "chrg_8b3g63gq2f"
//	}
//	if e := client.Do(dispute, createDispute); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute: %#v\n", dispute)
//
type CreateDispute struct {
	ChargeID string `json:"-"`
}

// Describe
func (req *CreateDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/disputes",
		ContentType: "application/json",
	}
}

