package omise

// ChargeStatus represents an enumeration of possible values for Charge.
type ChargeStatus string

// ChargeStatus can be one of the following list of constants:
const (
	ChargeExpired ChargeStatus = "expired"
	ChargeFailed ChargeStatus = "failed"
	ChargePending ChargeStatus = "pending"
	ChargeReversed ChargeStatus = "reversed"
	ChargeSuccessful ChargeStatus = "successful"
	ChargeUnknown ChargeStatus = "unknown"
)

