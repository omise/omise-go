package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	chain, list := &omise.Chain{}, &ListChains{
//	}
//	if e := client.Do(chain, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("chain: %#v\n", chain)
//
type ListChains struct {
	List
}

// Describe
func (req *ListChains) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/chains",
		ContentType: "application/json",
	}
}

// Example:
//
//	chain, retrieve := &omise.Chain{}, &RetrieveChain{
//		ChainID: "acch_8b3g64h9cm"
//	}
//	if e := client.Do(chain, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("chain: %#v\n", chain)
//
type RetrieveChain struct {
	ChainID string `json:"-"`
}

// Describe
func (req *RetrieveChain) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/chains/" + req.ChainID,
		ContentType: "application/json",
	}
}

// Example:
//
//	chain, revoke := &omise.Chain{}, &RevokeChain{
//		ChainID: "acch_8b3g64h9cm"
//	}
//	if e := client.Do(chain, revoke); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("chain: %#v\n", chain)
//
type RevokeChain struct {
	ChainID string `json:"-"`
}

// Describe
func (req *RevokeChain) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/chains/" + req.ChainID + "/revoke",
		ContentType: "application/json",
	}
}

