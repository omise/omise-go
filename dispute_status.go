package omise

// DisputeStatus represents an enumeration of possible status of a Dispute object, which
// can be one of the following list of constants:
type DisputeStatus string

const (
	Open    DisputeStatus = "open"
	Pending               = "pending"
	Closed                = "closed"
)
