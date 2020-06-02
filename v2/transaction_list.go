package omise

// TransactionList represents the list structure returned by Omise's REST API that contains
// Transaction struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type TransactionList struct {
	List
	Data []*Transaction `json:"data"`
}

// Find finds and returns Transaction with the given id. Returns nil if not found.
func (list *TransactionList) Find(id string) *Transaction {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}