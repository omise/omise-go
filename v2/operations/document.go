package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	document, destroy := &omise.Document{}, &DestroyDocument{
//		DisputeID: "dspt_9c4h4hr1eo"
//		DocumentID: "docu_9cdn8b6jt8"
//	}
//	if e := client.Do(document, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("document: %#v\n", document)
//
type DestroyDocument struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

// Describe
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
//	document, retrieve := &omise.Document{}, &RetrieveDocument{
//		DisputeID: "dspt_9c4h4hr1eo"
//		DocumentID: "docu_9cdn8b6jt8"
//	}
//	if e := client.Do(document, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("document: %#v\n", document)
//
type RetrieveDocument struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

// Describe
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
//	document, list := &omise.Document{}, &ListDocuments{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(document, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("document: %#v\n", document)
//
type ListDocuments struct {
	DisputeID string `json:"-"`
	List
}

// Describe
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
//	document, create := &omise.Document{}, &CreateDocument{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(document, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("document: %#v\n", document)
//
type CreateDocument struct {
	DisputeID string `json:"-"`
	File string `json:"file,omitempty"`
}

// Describe
func (req *CreateDocument) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/disputes/" + req.DisputeID + "/documents",
		ContentType: "application/json",
	}
}

