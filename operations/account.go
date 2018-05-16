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

func (req *RetrieveAccount) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/account",
		ContentType: "application/json",
	}
}
