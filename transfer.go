package omise

type Transfer struct {
	Base
	Recipient      string       `json:"recipient"`
	BankAccount    *BankAccount `json:"bank_account"`
	Sent           bool         `json:"sent"`
	Paid           bool         `json:"paid"`
	Amount         int64        `json:"amount"`
	Currency       string       `json:"currency"`
	FailureCode    *string      `json:"failure_code"`
	FailureMessage *string      `json:"failure_message"`
	Transaction    *string      `json:"transaction"`
}
