package omise

// DisputeStatus represents an enumeration of possible values for Dispute.
type DisputeStatus string

// DisputeStatus can be one of the following list of constants:
const (
	DisputeLost DisputeStatus = "lost"
	DisputeOpen DisputeStatus = "open"
	DisputePending DisputeStatus = "pending"
	DisputeWon DisputeStatus = "won"
)

