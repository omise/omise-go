package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Forex represents Forex object.
// See https://www.omise.co/forex-api for more information.
type Forex struct {
	BaseModel
	Base string `json:"base"`
	Location string `json:"location"`
	Quote string `json:"quote"`
	Rate float64 `json:"rate"`
}

// ForexService represents resource service.
type ForexService struct {
	*Client
}

// Forex defines resource service.
func (c *Client) Forex() *ForexService {
	return &ForexService{c}
}

// Retrieve retrieves forex
//
// Example:
//
//	forex, err := client.Forex().Retrieve(ctx, &RetrieveForexParams{
//		Currency: "undefined"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("forex: %#v\n", forex)
//
func (s *ForexService) Retrieve(ctx context.Context, params *RetrieveForexParams) (*Forex, error) {
	result := &Forex{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveForexParams params object.
type RetrieveForexParams struct {
	Currency string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveForexParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/forex/" + req.Currency,
		ContentType: "application/json",
	}
}

