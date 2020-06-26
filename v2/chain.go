package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Chain represents Chain object.
// See https://www.omise.co/chains-api for more information.
type Chain struct {
	Base
	Email string `json:"email"`
	Key string `json:"key"`
	Location string `json:"location"`
	Revoked bool `json:"revoked"`
}

// ChainService represents resource service.
type ChainService struct {
	*Client
}

// Chain defines resource service.
func (c *Client) Chain() *ChainService {
	return &ChainService{c}
}

// List lists chains
//
// Example:
//
//	chain, list := &omise.ChainList{}, &ListChains{
//	}
//	if e := client.Do(chain, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("chain: %#v\n", chain)
//
func (s *ChainService) List(ctx context.Context, params *ListChainsParams) (*ChainList, error) {
	result := &ChainList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListChainsParams params object.
type ListChainsParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListChainsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/chains",
		ContentType: "application/json",
	}
}

// Retrieve retrieves chain
//
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
func (s *ChainService) Retrieve(ctx context.Context, params *RetrieveChainParams) (*Chain, error) {
	result := &Chain{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveChainParams params object.
type RetrieveChainParams struct {
	ChainID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveChainParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/chains/" + req.ChainID,
		ContentType: "application/json",
	}
}

// Revoke revokes chain
//
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
func (s *ChainService) Revoke(ctx context.Context, params *RevokeChainParams) (*Chain, error) {
	result := &Chain{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RevokeChainParams params object.
type RevokeChainParams struct {
	ChainID string `json:"-"`
}

// Describe describes structure of request
func (req *RevokeChainParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/chains/" + req.ChainID + "/revoke",
		ContentType: "application/json",
	}
}

