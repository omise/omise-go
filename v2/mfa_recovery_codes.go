package omise

type MfaRecoveryCodes struct {
	Base
	Codes []interface{} `json:"codes"`
}
