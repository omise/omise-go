package operations

import (
	"net/url"

	"github.com/omise/omise-go/internal"
)

// Example:
//
//	charges, list := &omise.ChargeList{}, &ListCharges{
//		List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(charges, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of charges made in the last hour:", len(charges.Data))
//
type ListCharges struct {
	List
}

func (req *ListCharges) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges", nil}
}

// Note that because bool defaults to false in GO, we use DontCapture instead of Capture
// here so it matches with Omise's REST API default capture=true.
//
// Example:
//
//	charge, create := &omise.Charge{}, &CreateCharge{
//		Amount:      204842,
//		Currency:    "thb",
//		Description: "initial charge.",
//		Card:        token.ID,
//	}
//	if e := client.Do(charge, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created charge:", charge.ID)
//
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

// Example:
//
//	charge, update := &omise.Charge{}, &UpdateCharge{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateCharge struct {
	ChargeID    string `query:"-"`
	Description string
}

func (req *UpdateCharge) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/charges/" + req.ChargeID, nil}
}

// Example:
//
//	charge := &omise.Charge{ID: "chrg_323"}
//	if e := client.Do(charge, &RetrieveCharge{charge.ID}); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge #chrg_323: %#v\n", charge)
//
type RetrieveCharge struct {
	ChargeID string `query:"-"`
}

func (req *RetrieveCharge) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/charges/" + req.ChargeID, nil}
}

// If you have created a charge and passed capture=false you'll have an authorized only
// charge that you can capture at a later time. You can hold it for as long as permitted
// by the issuing bank. This delay may vary between cards from 1 to 30 days.
//
// Example:
//
//	charge := &omise.Charge{ID: "chrg_1234"}
//	if e := client.Do(charge, &CaptureCharge{charge.ID}); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("captured:", charge.Captured)
//
type CaptureCharge struct {
	ChargeID string
}

func (req *CaptureCharge) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/charges/" + req.ChargeID + "/capture", nil}
}
