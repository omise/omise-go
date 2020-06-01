package omise

// MfaRecoveryCodes represents resource object.
type MfaRecoveryCodes struct {
	Base
	Codes []interface{} `json:"codes"`
}

