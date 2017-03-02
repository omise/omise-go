package omise

// DisputeStatus represents an enumeration of possible status of a Dispute object.
type DisputeStatus string

// DisputeStatus can be one of the following list of constants:
const (
	Open    DisputeStatus = "open"
	Pending DisputeStatus = "pending"
	Won     DisputeStatus = "won"
	Lost    DisputeStatus = "lost"
	Closed  DisputeStatus = "closed" // meta-status only for querying, does not actually appear.
)
