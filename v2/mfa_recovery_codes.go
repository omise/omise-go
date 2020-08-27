package omise

// MfaRecoveryCodes represents resource object.
type MfaRecoveryCodes struct {
	Base
	Codes []interface{} `json:"codes"`
}

// MfaRecoveryCodesService represents resource service.
type MfaRecoveryCodesService struct {
	*Client
}

// MfaRecoveryCodes defines resource service.
func (c *Client) MfaRecoveryCodes() *MfaRecoveryCodesService {
	return &MfaRecoveryCodesService{c}
}

