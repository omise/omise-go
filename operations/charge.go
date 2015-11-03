package operations

import (
	"github.com/omise/omise-go/internal"
	"net/url"
)

type CreateCharge struct {
	Customer    string
	Card        string
	Amount      int64
	Currency    string
	Description string
	DontCapture bool   `query:"-"`
	ReturnURI   string `query:"return_uri"`
}

func (req *CreateCharge) Op() *internal.Op {
	op := &internal.Op{internal.API, "POST", "/charges", url.Values{}}
	if req.DontCapture {
		op.Values.Set("capture", "false") // defaults to true
	}
	return op
}
