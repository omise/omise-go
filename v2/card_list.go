package omise

// CardList represents the list structure returned by Omise's REST API that contains
// Card struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type CardList struct {
	List
	Data []*Card `json:"data"`
}

// Find finds and returns Card with the given id. Returns nil if not found.
func (list *CardList) Find(id string) *Card {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}