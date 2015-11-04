package operations

import (
	"net/url"

	"github.com/omise/omise-go/internal"
)

type ListCharges struct {
	List
}

func (req *ListCharges) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges", nil}
}

type CreateCharge struct {
	Customer    string
	Card        string
	Amount      int64
	Currency    string
	Description string
	DontCapture bool   `query:"-"` // inverse, since `capture` defaults to true
	ReturnURI   string `query:"return_uri"`
}

func (req *CreateCharge) Op() *internal.Op {
	op := &internal.Op{internal.API, "POST", "/charges", url.Values{}}
	if req.DontCapture {
		op.Values.Set("capture", "false")
	}
	return op
}

type UpdateCharge struct {
	ChargeID    string `query:"-"`
	Description string
}

func (req *UpdateCharge) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/charges/" + req.ChargeID, nil}
}

type RetreiveCharge struct {
	ChargeID string `query:"-"`
}

func (req *RetreiveCharge) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges/" + req.ChargeID, nil}
}

type CaptureCharge struct {
	ChargeID string
}

func (req *CaptureCharge) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/charges/" + req.ChargeID + "/capture", nil}
}
