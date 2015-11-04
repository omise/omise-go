package omise

type Transaction struct {
	Base
	Amount   int64           `json:"amount"`
	Currency string          `json:"currency"`
	Type     TransactionType `json:"type"`
}

type TransactionList struct {
	List
	Data []*Transaction `json:"data"`
}

func (list *TransactionList) Find(id string) *Transaction {
	for _, tx := range list.Data {
		if tx.ID == id {
			return tx
		}
	}

	return nil
}
