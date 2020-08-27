package omise

// LinkList represents the list structure returned by Omise's REST API that contains
// Link struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type LinkList struct {
	List
	Data []*Link `json:"data"`
}

// Find finds and returns Link with the given id. Returns nil if not found.
func (list *LinkList) Find(id string) *Link {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
