package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	balance, retrieve := &omise.Balance{}, &RetrieveBalance{
//	}
//	if e := client.Do(balance, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("balance: %#v\n", balance)
//
type RetrieveBalance struct {
}

// Describe
func (req *RetrieveBalance) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/balance",
		ContentType: "application/json",
	}
}

