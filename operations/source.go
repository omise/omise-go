package operations

import (
	"encoding/json"

	"github.com/omise/omise-go/internal"
)

// Example: Create source with public key
//
//	source, createSource := &omise.Source{}, &CreateSourceWithPublicKey{
//		Amount:   100000,
//		Currency: "thb",
//		Type:     "bill_payment_tesco_lotus",
//	}
//	if e := client.Do(source, createSource); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created source:", source.ID)
//
type CreateSourceWithPublicKey struct {
	Type     string
	Amount   int64
	Currency string
}

func (req *CreateSourceWithPublicKey) MarshalJSON() ([]byte, error) {
	param := struct {
		Type     string `json:"type"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
	}{
		Type:     req.Type,
		Amount:   req.Amount,
		Currency: req.Currency,
	}

	return json.Marshal(param)
}

func (req *CreateSourceWithPublicKey) Op() *internal.Op {
	return &internal.Op{
		Endpoint:          internal.API,
		Method:            "POST",
		Path:              "/sources",
		ContentType:       "application/json",
		ForceUsePublicKey: true,
	}
}

// Example: Create source with secret key
//
//	source, createSource := &omise.Source{}, &CreateSourceWithSecretKey{
//		Amount:   100000,
//		Currency: "thb",
//		Type:     "bill_payment_tesco_lotus",
//	}
//	if e := client.Do(source, createSource); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created source:", source.ID)
type CreateSourceWithSecretKey struct {
	Type     string `json:"type"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func (req *CreateSourceWithSecretKey) MarshalJSON() ([]byte, error) {
	param := struct {
		Type     string `json:"type"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
	}{
		Type:     req.Type,
		Amount:   req.Amount,
		Currency: req.Currency,
	}

	return json.Marshal(param)
}

func (req *CreateSourceWithSecretKey) Op() *internal.Op {
	return &internal.Op{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/sources",
		ContentType: "application/json",
	}
}

// Example: Retrieve Source
//
//	source, retrieve := &omise.Source{}, &RetrieveSource{"src_123"}
//	if e := client.Do(source, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("source #123: %#v\n", source)
//
type RetrieveSource struct {
	SourceID string `query:"-"`
}

func (req *RetrieveSource) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/sources/" + req.SourceID,
	}
}
