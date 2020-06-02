package omise

// RefundList represents the list structure returned by Omise's REST API that contains
// Refund struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type RefundList struct {
	List
	Data []*Refund `json:"data"`
}

// Find finds and returns Refund with the given id. Returns nil if not found.
func (list *RefundList) Find(id string) *Refund {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}