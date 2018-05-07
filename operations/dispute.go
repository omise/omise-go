package operations

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	disputes, list := &omise.DisputeList{}, &ListDisputes{
//		State: omise.Open,
//	}
//	if e := client.Do(disputes, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of open disputes:", len(disputes.Data))
//
type ListDisputes struct {
	State omise.DisputeStatus
	List
}

func (req *ListDisputes) Describe() *internal.Description {
	path := "/disputes"
	switch req.State {
	case omise.Open:
		path += "/open"
	case omise.Pending:
		path += "/pending"
	case omise.Won, omise.Lost, omise.Closed:
		path += "/closed"
	}

	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        path,
		ContentType: "application/json",
	}
}

// Example:
//
//	dispute, retrieve := &omise.Dispute{}, &RetrieveDispute{"dspt_123"}
//	if e := client.Do(dispute, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("dispute #123: %#v\n", dispute)
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
//	dispute, update := &omise.Dispute{}, &UpdateDispute{
//		DisputeID: "dspt_777",
//		Message:   "update me!",
//	}
//	if e := client.Do(dispute, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated dispute: %#v\n", dispute)
//
type UpdateDispute struct {
	DisputeID string                 `json:"-"`
	Message   string                 `json:"message"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

func (req *UpdateDispute) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/disputes/" + req.DisputeID,
		ContentType: "application/json",
	}
}
