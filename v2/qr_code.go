package omise

// QrCode represents resource object.
type QrCode struct {
	Base
	Email string `json:"email"`
	MfaProvisioningURI string `json:"mfa_provisioning_uri"`
	Secret string `json:"secret"`
}

