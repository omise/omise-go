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
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges",
		ContentType: "application/json",
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
	Customer    string                 `json:"customer,omitempty"`
	Card        string                 `json:"card,omitempty"`
	Source      string                 `json:"source,omitempty"`
	Amount      int64                  `json:"amount"`
	Currency    string                 `json:"currency"`
	Offsite     omise.OffsiteTypes     `json:"offsite,omitempty"`
	Description string                 `json:"description,omitempty"`
	DontCapture bool                   `json:"-"` // inverse, since `capture` defaults to true
	ReturnURI   string                 `json:"return_uri,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func (req *CreateCharge) MarshalJSON() ([]byte, error) {
	type Alias CreateCharge
	params := struct {
		*Alias
		Capture *bool `json:"capture,omitempty"`
	}{
		Alias: (*Alias)(req),
	}
	if params.DontCapture {
		params.Capture = new(bool)
	}
	return json.Marshal(params)
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
	ChargeID    string                 `json:"-"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func (req *UpdateCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
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
	ChargeID string `json:"-"`
}

func (req *RetrieveCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/charges/" + req.ChargeID,
		ContentType: "application/json",
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
	ChargeID string `json:"-"`
}

func (req *CaptureCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/capture",
		ContentType: "application/json",
	}
}

// If you have created a charge and passed capture=false you'll have an authorized only
// charge that can be reversed, releasing held money, at a later time without incurring a
// refund fee.
type ReverseCharge struct {
	ChargeID string `json:"-"`
}

func (req *ReverseCharge) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/charges/" + req.ChargeID + "/reverse",
		ContentType: "application/json",
	}
}
