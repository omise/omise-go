package omise

// DisputeReasonCode represents an enumeration of possible values for Dispute.
type DisputeReasonCode string

// DisputeReasonCode can be one of the following list of constants:
const (
	DisputeCancelledRecurringTransaction DisputeReasonCode = "cancelled_recurring_transaction"
	DisputeCreditNotProcessed DisputeReasonCode = "credit_not_processed"
	DisputeDuplicateProcessing DisputeReasonCode = "duplicate_processing"
	DisputeExpiredCard DisputeReasonCode = "expired_card"
	DisputeGoodsOrServicesNotProvided DisputeReasonCode = "goods_or_services_not_provided"
	DisputeIncorrectCurrency DisputeReasonCode = "incorrect_currency"
	DisputeIncorrectTransactionAmount DisputeReasonCode = "incorrect_transaction_amount"
	DisputeLatePresentment DisputeReasonCode = "late_presentment"
	DisputeNonMatchingAccountNumber DisputeReasonCode = "non_matching_account_number"
	DisputeNotAsDescribedOrDefectiveMerchandise DisputeReasonCode = "not_as_described_or_defective_merchandise"
	DisputeNotAvailable DisputeReasonCode = "not_available"
	DisputeNotRecorded DisputeReasonCode = "not_recorded"
	DisputeOther DisputeReasonCode = "other"
	DisputePaidByOtherMeans DisputeReasonCode = "paid_by_other_means"
	DisputeTransactionNotRecognised DisputeReasonCode = "transaction_not_recognised"
	DisputeUnauthorizedCharge DisputeReasonCode = "unauthorized_charge"
)

