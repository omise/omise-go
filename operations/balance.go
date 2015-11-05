package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	balance := &omise.Balance{}
//	if e := client.Do(balance, &RetreiveBalance{}); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("we have $", balance.Available)
//
type RetreiveBalance struct{}

func (req *RetreiveBalance) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/balance", nil}
}
