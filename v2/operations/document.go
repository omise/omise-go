package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Document{}, &DestroyDocument{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyDocument struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

func (req *DestroyDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Document{}, &RetrieveDocument{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveDocument struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

func (req *RetrieveDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Document{}, &ListDocuments{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListDocuments struct {
	DisputeID string `json:"-"`
}

func (req *ListDocuments) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID + "/documents",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Document{}, &CreateDocument{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateDocument struct {
	DisputeID string `json:"-"`
	File string `json:"file"`
}

func (req *CreateDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/disputes/" + req.DisputeID + "/documents",
		ContentType: "application/json",
	}
}

