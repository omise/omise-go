package omise

type DisputeStatus string

const (
	DisputeLost DisputeStatus = "lost"
	DisputeOpen DisputeStatus = "open"
	DisputePending DisputeStatus = "pending"
	DisputeWon DisputeStatus = "won"
)
