package omise

// Transfer represents Omise's transfer object.
// See https://www.omise.co/transfers-api for more information.
type Transfer struct {
	Base
	Recipient   string       `json:"recipient" pretty:""`
	BankAccount *BankAccount `json:"bank_account"`

	Sent     bool   `json:"sent" pretty:""`
	Paid     bool   `json:"paid" pretty:""`
	Amount   int64  `json:"amount" pretty:""`
	Currency string `json:"currency" pretty:""`

	FailureCode    *string `json:"failure_code"`
	FailureMessage *string `json:"failure_message"`
	Transaction    *string `json:"transaction"`
}
