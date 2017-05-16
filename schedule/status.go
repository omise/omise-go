package schedule

// Status represents an enumeration of possible status of a Schedule object.
type Status string

// Status can be one of the following list of constants:
const (
	Active    Status = "active"
	Expiring  Status = "expiring"
	Expired   Status = "expired"
	Deleted   Status = "deleted"
	Suspended Status = "suspended"
)
