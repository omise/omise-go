package omise

// ChargeList represents the list structure returned by Omise's REST API that contains
// Charge struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type ChargeList struct {
	List
	Data []*Charge `json:"data"`
}

// Find finds and returns Charge with the given id. Returns nil if not found.
func (list *ChargeList) Find(id string) *Charge {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}