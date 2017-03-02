package omise

// SourceOfFunds represents an enumeration of possible source of funds on a Charge object.
type SourceOfFunds string

// SourceOfFunds can be one the following list of constants:
const (
	FromCard    SourceOfFunds = "card"
	FromOffsite SourceOfFunds = "offsite"
)
