package omise

// Transfer represents Omise's transfer object.
// See https://www.omise.co/transfers-api for more information.
type Transfer struct {
	Base
	Recipient   string       `json:"recipient"`
	BankAccount *BankAccount `json:"bank_account"`

	Sent     bool   `json:"sent"`
	Paid     bool   `json:"paid"`
	Fee      int64  `json:"fee"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`

	FailureCode    *string `json:"failure_code"`
	FailureMessage *string `json:"failure_message"`
	Transaction    *string `json:"transaction"`

	Metadata map[string]interface{} `json:"metadata"`
}
