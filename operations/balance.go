package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	balance := &omise.Balance{}
//	if e := client.Do(balance, &RetrieveBalance{}); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("we have $", balance.Available)
//
type RetrieveBalance struct{}

func (req *RetrieveBalance) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/balance",
		ContentType: "application/json",
	}
}
