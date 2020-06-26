package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Balance represents Balance object.
// See https://www.omise.co/balance-api for more information.
type Balance struct {
	Base
	Currency string `json:"currency"`
	Location string `json:"location"`
	Reserve int64 `json:"reserve"`
	Total int64 `json:"total"`
	Transferable int64 `json:"transferable"`
}

// BalanceService represents resource service.
type BalanceService struct {
	*Client
}

// Balance defines resource service.
func (c *Client) Balance() *BalanceService {
	return &BalanceService{c}
}

// Retrieve retrieves balance
//
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
func (s *BalanceService) Retrieve(ctx context.Context, params *RetrieveBalanceParams) (*Balance, error) {
	result := &Balance{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveBalanceParams params object.
type RetrieveBalanceParams struct {
}

// Describe describes structure of request
func (req *RetrieveBalanceParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/balance",
		ContentType: "application/json",
	}
}

