package omise

// ChargeStatus represents an enumeration of possible status of a Charge object.
type ChargeStatus string

// ChargeStatus can be one of the following list of constants:
const (
	ChargeFailed     ChargeStatus = "failed"
	ChargePending    ChargeStatus = "pending"
	ChargeSuccessful ChargeStatus = "successful"
	ChargeReversed   ChargeStatus = "reversed"
)
