package omise

type AuthenticationType string

const (
	ThreeDS AuthenticationType = "3DS"
	Passkey AuthenticationType = "PASSKEY"
)
