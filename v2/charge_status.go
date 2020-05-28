package omise

type ChargeStatus string

const (
	ChargeExpired ChargeStatus = "expired"
	ChargeFailed ChargeStatus = "failed"
	ChargePending ChargeStatus = "pending"
	ChargeReversed ChargeStatus = "reversed"
	ChargeSuccessful ChargeStatus = "successful"
)
