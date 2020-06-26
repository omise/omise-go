package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Document represents Document object.
// See https://www.omise.co/documents-api for more information.
type Document struct {
	Base
	DownloadURI string `json:"download_uri"`
	Filename string `json:"filename"`
	Location string `json:"location"`
}

// DocumentService represents resource service.
type DocumentService struct {
	*Client
}

// Document defines resource service.
func (c *Client) Document() *DocumentService {
	return &DocumentService{c}
}

// Destroy destroys document
//
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
func (s *DocumentService) Destroy(ctx context.Context, params *DestroyDocumentParams) (*Document, error) {
	result := &Document{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DestroyDocumentParams params object.
type DestroyDocumentParams struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

// Describe describes structure of request
func (req *DestroyDocumentParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
		ContentType: "application/json",
	}
}

// Retrieve retrieves document
//
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
func (s *DocumentService) Retrieve(ctx context.Context, params *RetrieveDocumentParams) (*Document, error) {
	result := &Document{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveDocumentParams params object.
type RetrieveDocumentParams struct {
	DisputeID string `json:"-"`
	DocumentID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveDocumentParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
		ContentType: "application/json",
	}
}

// List lists documents
//
// Example:
//
//	document, list := &omise.DocumentList{}, &ListDocuments{
//		DisputeID: "dspt_9c4h4hr1eo"
//	}
//	if e := client.Do(document, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("document: %#v\n", document)
//
func (s *DocumentService) List(ctx context.Context, params *ListDocumentsParams) (*DocumentList, error) {
	result := &DocumentList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListDocumentsParams params object.
type ListDocumentsParams struct {
	DisputeID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListDocumentsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/disputes/" + req.DisputeID + "/documents",
		ContentType: "application/json",
	}
}

// Create creates document
//
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
func (s *DocumentService) Create(ctx context.Context, params *CreateDocumentParams) (*Document, error) {
	result := &Document{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateDocumentParams params object.
type CreateDocumentParams struct {
	DisputeID string `json:"-"`
	File string `json:"file,omitempty"`
}

// Describe describes structure of request
func (req *CreateDocumentParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/disputes/" + req.DisputeID + "/documents",
		ContentType: "application/json",
	}
}

