package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	link, destroy := &omise.Link{}, &DestroyLink{
//		LinkID: "link_7ak4h9cm6j"
//	}
//	if e := client.Do(link, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
type DestroyLink struct {
	LinkID string `json:"-"`
}

// Describe
func (req *DestroyLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}

// Example:
//
//	link, retrieve := &omise.Link{}, &RetrieveLink{
//		LinkID: "link_7ak4h9cm6j"
//	}
//	if e := client.Do(link, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
type RetrieveLink struct {
	LinkID string `json:"-"`
}

// Describe
func (req *RetrieveLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}

// Example:
//
//	link, list := &omise.Link{}, &ListLinks{
//	}
//	if e := client.Do(link, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
type ListLinks struct {
	List
}

// Describe
func (req *ListLinks) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links",
		ContentType: "application/json",
	}
}

// Example:
//
//	link, create := &omise.Link{}, &CreateLink{
//		Amount: 10000
//		Currency: "thb"
//		Description: "undefined"
//		Title: "undefined"
//	}
//	if e := client.Do(link, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("link: %#v\n", link)
//
type CreateLink struct {
	Amount int64 `json:"amount"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Multiple bool `json:"multiple,omitempty"`
	Title string `json:"title"`
}

// Describe
func (req *CreateLink) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/links",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, listCharges := &omise.Charge{}, &ListLinkCharges{
//		LinkID: "link_7ak4h9cm6j"
//	}
//	if e := client.Do(charge, listCharges); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("charge: %#v\n", charge)
//
type ListLinkCharges struct {
	LinkID string `json:"-"`
	List
}

// Describe
func (req *ListLinkCharges) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/links/" + req.LinkID + "/charges",
		ContentType: "application/json",
	}
}

