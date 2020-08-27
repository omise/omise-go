package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

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
type DestroyTransfer struct {
	TransferID string `json:"-"`
}

// Describe
func (req *DestroyTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

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
type RetrieveTransfer struct {
	TransferID string `json:"-"`
}

// Describe
func (req *RetrieveTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

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
type UpdateTransfer struct {
	TransferID string `json:"-"`
	Amount int64 `json:"amount,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe
func (req *UpdateTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

// Example:
//
//	transfer, list := &omise.Transfer{}, &ListTransfers{
//	}
//	if e := client.Do(transfer, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer: %#v\n", transfer)
//
type ListTransfers struct {
	List
}

// Describe
func (req *ListTransfers) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transfers",
		ContentType: "application/json",
	}
}

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
type CreateTransfer struct {
	Amount int64 `json:"amount,omitempty"`
	FailFast bool `json:"fail_fast,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Recipient string `json:"recipient,omitempty"`
}

// Describe
func (req *CreateTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/transfers",
		ContentType: "application/json",
	}
}

// Example:
//
//	schedule, listSchedules := &omise.Schedule{}, &ListTransferSchedules{
//	}
//	if e := client.Do(schedule, listSchedules); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
type ListTransferSchedules struct {
	List
}

// Describe
func (req *ListTransferSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transfers/schedules",
		ContentType: "application/json",
	}
}

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
type MarkTransferAsPaid struct {
	TransferID string `json:"-"`
}

// Describe
func (req *MarkTransferAsPaid) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/transfers/" + req.TransferID + "/mark_as_paid",
		ContentType: "application/json",
	}
}

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
type MarkTransferAsSent struct {
	TransferID string `json:"-"`
}

// Describe
func (req *MarkTransferAsSent) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/transfers/" + req.TransferID + "/mark_as_sent",
		ContentType: "application/json",
	}
}

