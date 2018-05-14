package operations

import (
	"encoding/json"

	"github.com/omise/omise-go"
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

func (req *ListCharges) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/charges",
	}
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
	Source      string
	Amount      int64
	Currency    string
	Offsite     omise.OffsiteTypes
	Description string
	DontCapture bool   `query:"-"` // inverse, since `capture` defaults to true
	ReturnURI   string `query:"return_uri"`
}

func (req *CreateCharge) MarshalJSON() ([]byte, error) {
	param := map[string]interface{}{
		"amount":   req.Amount,
		"currency": req.Currency,
	}

	if req.Customer != "" {
		param["customer"] = req.Customer
	}

	if req.Card != "" {
		param["card"] = req.Card
	}

	if req.Source != "" {
		param["source"] = req.Source
	}

	if req.Offsite != "" {
		param["offsite"] = req.Offsite
	}

	if req.ReturnURI != "" {
		param["return_uri"] = req.ReturnURI
	}

	if req.DontCapture {
		param["capture"] = false
	}

	return json.Marshal(param)
}

func (req *CreateCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges",
		ContentType: "application/json",
	}
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

func (req *UpdateCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "PATCH",
		Path:     "/charges/" + req.ChargeID,
	}
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

func (req *RetrieveCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/charges/" + req.ChargeID,
	}
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

func (req *CaptureCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "POST",
		Path:     "/charges/" + req.ChargeID + "/capture",
	}
}

// If you have created a charge and passed capture=false you'll have an authorized only
// charge that can be reversed, releasing held money, at a later time without incurring a
// refund fee.
type ReverseCharge struct {
	ChargeID string
}

func (req *ReverseCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint: internal.API,
		Method:   "POST",
		Path:     "/charges/" + req.ChargeID + "/reverse",
	}
}
