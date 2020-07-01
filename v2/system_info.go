package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// SystemInfo represents resource object.
type SystemInfo struct {
	Base
	Location string `json:"location"`
	Versions []string `json:"versions"`
}

// SystemInfoService represents resource service.
type SystemInfoService struct {
	*Client
}

// SystemInfo defines resource service.
func (c *Client) SystemInfo() *SystemInfoService {
	return &SystemInfoService{c}
}

// Retrieve retrieves system info
//
// Example:
//
//	system_info, err := client.SystemInfo().Retrieve(ctx, &RetrieveSystemInfoParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("system_info: %#v\n", system_info)
//
func (s *SystemInfoService) Retrieve(ctx context.Context, params *RetrieveSystemInfoParams) (*SystemInfo, error) {
	result := &SystemInfo{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveSystemInfoParams params object.
type RetrieveSystemInfoParams struct {
}

// Describe describes structure of request
func (req *RetrieveSystemInfoParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/",
		ContentType: "application/json",
	}
}

