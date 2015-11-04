package omise

type DisputeStatus string

const (
	Any     DisputeStatus = "" // for search/test only
	Open                  = "open"
	Pending               = "pending"
	Closed                = "closed"
)
