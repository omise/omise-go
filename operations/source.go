package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example: Create source
//
//	source, createSource := &omise.Source{}, &CreateSource{
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
type CreateSource struct {
	Type     string `json:"type"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func (req *CreateSource) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/sources",
		ContentType: "application/json",
		APIKey:      "public",
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
	SourceID string `json:"-"`
}

func (req *RetrieveSource) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/sources/" + req.SourceID,
		ContentType: "application/json",
	}
}
