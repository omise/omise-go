package operations

import (
	"github.com/omise/omise-go/internal"
	"io"
)

type ListDisputeDocuments struct {
	List
	DisputeID string `query:"-"`
}

func (req *ListDisputeDocuments) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "disputes/" + req.DisputeID + "/documents",
	}
}

type UploadDisputeDocument struct {
	DisputeID string    `query:"-"`
	Content   io.Reader // TODO: File struct
}

func (req *UploadDisputeDocument) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "POST",
		Path:     "disputes/" + req.DisputeID + "/documents",
	}
}

type RetrieveDisputeDocument struct {
	DisputeID  string `query:"-"`
	DocumentID string `query:"-"`
}

func (req *RetrieveDisputeDocument) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
	}
}

type DestroyDisputeDocument struct {
	DisputeID  string `query:"-"`
	DocumentID string `query:"-"`
}

func (req *DestroyDisputeDocument) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "DELETE",
		Path:     "disputes/" + req.DisputeID + "/documents/" + req.DocumentID,
	}
}
