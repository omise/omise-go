package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	forex, retrieve := &omise.Forex{}, &RetrieveForex{
//		Currency: "undefined"
//	}
//	if e := client.Do(forex, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("forex: %#v\n", forex)
//
type RetrieveForex struct {
	Currency string `json:"-"`
}

// Describe
func (req *RetrieveForex) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/forex/" + req.Currency,
		ContentType: "application/json",
	}
}

