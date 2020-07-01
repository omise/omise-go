package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Link represents Links object.
// See https://www.omise.co/links-api for more information.
type Link struct {
	Base
	Amount int64 `json:"amount"`
	Charges *ChargeList `json:"charges"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Location string `json:"location"`
	Multiple bool `json:"multiple"`
	PaymentURI string `json:"payment_uri"`
	Title string `json:"title"`
	Used bool `json:"used"`
	UsedAt time.Time `json:"used_at"`
}

// LinkService represents resource service.
type LinkService struct {
	*Client
}

// Link defines resource service.
func (c *Client) Link() *LinkService {
	return &LinkService{c}
}

// Destroy destroys link
//
// Example:
//
//	link, err := client.Link().Destroy(ctx, &DestroyLinkParams{
//		LinkID: "link_7ak4h9cm6j"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
func (s *LinkService) Destroy(ctx context.Context, params *DestroyLinkParams) (*Link, error) {
	result := &Link{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DestroyLinkParams params object.
type DestroyLinkParams struct {
	LinkID string `json:"-"`
}

// Describe describes structure of request
func (req *DestroyLinkParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}

// Retrieve retrieves link
//
// Example:
//
//	link, err := client.Link().Retrieve(ctx, &RetrieveLinkParams{
//		LinkID: "link_7ak4h9cm6j"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
func (s *LinkService) Retrieve(ctx context.Context, params *RetrieveLinkParams) (*Link, error) {
	result := &Link{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveLinkParams params object.
type RetrieveLinkParams struct {
	LinkID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveLinkParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}

// List lists links
//
// Example:
//
//	link, err := client.Link().List(ctx, &ListLinksParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
func (s *LinkService) List(ctx context.Context, params *ListLinksParams) (*LinkList, error) {
	result := &LinkList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListLinksParams params object.
type ListLinksParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListLinksParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links",
		ContentType: "application/json",
	}
}

// Create creates link
//
// Example:
//
//	link, err := client.Link().Create(ctx, &CreateLinkParams{
//		Amount: 10000
//		Currency: "thb"
//		Description: "undefined"
//		Title: "undefined"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
func (s *LinkService) Create(ctx context.Context, params *CreateLinkParams) (*Link, error) {
	result := &Link{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateLinkParams params object.
type CreateLinkParams struct {
	Amount int64 `json:"amount"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Multiple bool `json:"multiple,omitempty"`
	Title string `json:"title"`
}

// Describe describes structure of request
func (req *CreateLinkParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/links",
		ContentType: "application/json",
	}
}

// ListCharges lists link charges
//
// Example:
//
//	charge, err := client.Link().ListCharges(ctx, &ListLinkChargesParams{
//		LinkID: "link_7ak4h9cm6j"
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
func (s *LinkService) ListCharges(ctx context.Context, params *ListLinkChargesParams) (*ChargeList, error) {
	result := &ChargeList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListLinkChargesParams params object.
type ListLinkChargesParams struct {
	LinkID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListLinkChargesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links/" + req.LinkID + "/charges",
		ContentType: "application/json",
	}
}

