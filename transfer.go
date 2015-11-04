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

type TransferList struct {
	List
	Data []*Transfer `json:"data"`
}

func (list *TransferList) Find(id string) *Transfer {
	for _, transfer := range list.Data {
		if transfer.ID == id {
			return transfer
		}
	}

	return nil
}
