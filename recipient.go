package omise

// Recipient represents Omise's recipient object.
// See https://www.omise.co/recipients-api for more information.
type Recipient struct {
	Base
	Verified    bool          `json:"verified"`
	Active      bool          `json:"active"`
	Name        string        `json:"name" pretty:""`
	Email       string        `json:"email" pretty:""`
	Description *string       `json:"description"`
	Type        RecipientType `json:"type" pretty:""`
	TaxID       *string       `json:"tax_id"`
	BankAccount *BankAccount  `json:"bank_account"`
	FailureCode *string       `json:"failure_code"`
}
