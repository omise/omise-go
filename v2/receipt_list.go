package omise

// ReceiptList represents the list structure returned by Omise's REST API that contains
// Receipt struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type ReceiptList struct {
	List
	Data []*Receipt `json:"data"`
}

// Find finds and returns Receipt with the given id. Returns nil if not found.
func (list *ReceiptList) Find(id string) *Receipt {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}