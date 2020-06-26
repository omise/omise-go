package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Capability represents Capability object.
// See https://www.omise.co/capability-api for more information.
type Capability struct {
	Base
	Banks []string `json:"banks"`
	Country string `json:"country"`
	Location string `json:"location"`
	PaymentMethods []PaymentMethod `json:"payment_methods"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}

// CapabilityService represents resource service.
type CapabilityService struct {
	*Client
}

// Capability defines resource service.
func (c *Client) Capability() *CapabilityService {
	return &CapabilityService{c}
}

// Retrieve retrieves capability
//
// Example:
//
//	capability, retrieve := &omise.Capability{}, &RetrieveCapability{
//	}
//	if e := client.Do(capability, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("capability: %#v\n", capability)
//
func (s *CapabilityService) Retrieve(ctx context.Context, params *RetrieveCapabilityParams) (*Capability, error) {
	result := &Capability{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveCapabilityParams params object.
type RetrieveCapabilityParams struct {
}

// Describe describes structure of request
func (req *RetrieveCapabilityParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/capability",
		ContentType: "application/json",
	}
}

