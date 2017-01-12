package omise

// SourceOfFunds represents an enumeration of possible source of funds on a Charge object,
// which can be one of the following list of constants:
type SourceOfFunds string

const (
	FromCard    SourceOfFunds = "card"
	FromOffsite SourceOfFunds = "offsite"
)
