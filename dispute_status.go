package omise

type DisputeStatus string

const (
	Open    DisputeStatus = "open"
	Pending               = "pending"
	Closed                = "closed"
)
