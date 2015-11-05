package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	account := &omise.Account{}
//	if e := client.Do(account, &RetrieveAccount{}); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("my account!: %#v\n", account)
//
type RetrieveAccount struct{}

func (req *RetrieveAccount) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/account", nil}
}
