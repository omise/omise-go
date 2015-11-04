package operations

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal"
)

type ListDisputes struct {
	State omise.DisputeStatus
	List
}

func (req *ListDisputes) Op() *internal.Op {
	path := "/disputes"
	if req.State != omise.Any {
		path += "/" + string(req.State)
	}

	return &internal.Op{internal.API, "GET", path, nil}
}

type RetreiveDispute struct {
	DisputeID string `query:"-"`
}

func (req *RetreiveDispute) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/disputes/" + req.DisputeID, nil}
}

type UpdateDispute struct {
	DisputeID string `query:"-"`
	Message   string
}

func (req *UpdateDispute) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/disputes/" + req.DisputeID, nil}
}
