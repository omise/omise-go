package omise

// ChargeFailureCode represents an enumeration of possible values for Charge.
type ChargeFailureCode string

// ChargeFailureCode can be one of the following list of constants:
const (
	ChargeFailedFraudCheck ChargeFailureCode = "failed_fraud_check"
	ChargeFailedProcessing ChargeFailureCode = "failed_processing"
	ChargeInsufficientBalance ChargeFailureCode = "insufficient_balance"
	ChargeInsufficientFund ChargeFailureCode = "insufficient_fund"
	ChargeInvalidAccountNumber ChargeFailureCode = "invalid_account_number"
	ChargeInvalidSecurityCode ChargeFailureCode = "invalid_security_code"
	ChargePaymentCancelled ChargeFailureCode = "payment_cancelled"
	ChargePaymentRejected ChargeFailureCode = "payment_rejected"
	ChargeTimeout ChargeFailureCode = "timeout"
)

