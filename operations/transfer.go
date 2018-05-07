package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	transfers, list := &omise.TransferList{}, &ListTransfers{
//		List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(transfers, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of transfers in the last hour:", len(transfers.Data))
//
type ListTransfers struct {
	List
}

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
//		Amount: 32100,
//	}
//	if e := client.Do(transfer, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("transferred to default recipient with:", transfer.ID)
//
type CreateTransfer struct {
	Amount    int64                  `json:"amount,omitempty"`
	Recipient string                 `json:"recipient,omitempty"`
	FailFast  bool                   `json:"fail_fast,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

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
//	transfer, retrieve := &omise.Transfer{}, &RetrieveTransfer{"trsf_123"}
//	if e := client.Do(transfer, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transfer #123: %#v\n", transfer)
//
type RetrieveTransfer struct {
	TransferID string `json:"-"`
}

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
//		TransferID: "trsf_777",
//		Amount:     12300,
//	}
//	if e := client.Do(transfer, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated transfer #777: %#v\n", transfer)
//
type UpdateTransfer struct {
	TransferID string                 `json:"-"`
	Amount     int64                  `json:"amount,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

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
//	del, destroy := &omise.Deletion{}, &DestroyTransfer{"trsf_777"}
//	if e := client.Do(del, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("not transferring anymore:", del.ID)
//
type DestroyTransfer struct {
	TransferID string `json:"-"`
}

func (req *DestroyTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}
