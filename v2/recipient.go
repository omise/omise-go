package omise

import (
	"time"
)

// Recipient represents Recipient object.
// See https://www.omise.co/recipients-api for more information.
type Recipient struct {
	Base
	ActivatedAt time.Time `json:"activated_at"`
	Active bool `json:"active"`
	BankAccount *BankAccount `json:"bank_account"`
	Default bool `json:"default"`
	Description *string `json:"description"`
	Email string `json:"email"`
	FailureCode *RecipientFailureCode `json:"failure_code"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Name string `json:"name"`
	Schedule string `json:"schedule"`
	TaxID *string `json:"tax_id"`
	Type *RecipientType `json:"type"`
	Verified bool `json:"verified"`
	VerifiedAt time.Time `json:"verified_at"`
}