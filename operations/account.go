package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	account := &omise.Account{}
//	if e := client.Do(account, &RetreiveAccount{}); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("my account!: %#v\n", account)
//
type RetreiveAccount struct{}

func (req *RetreiveAccount) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/account", nil}
}
