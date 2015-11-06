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
	State omise.DisputeStatus `query:"-"`
	List
}

func (req *ListDisputes) Op() *internal.Op {
	path := "/disputes"
	switch req.State {
	case omise.Open:
		path += "/open"
	case omise.Pending:
		path += "/pending"
	case omise.Won, omise.Lost, omise.Closed:
		path += "/closed"
	}

	return &internal.Op{internal.API, "GET", path, nil}
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
	DisputeID string `query:"-"`
}

func (req *RetrieveDispute) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/disputes/" + req.DisputeID, nil}
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
	DisputeID string `query:"-"`
	Message   string
}

func (req *UpdateDispute) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/disputes/" + req.DisputeID, nil}
}
