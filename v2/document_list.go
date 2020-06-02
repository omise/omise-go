package omise

// DocumentList represents the list structure returned by Omise's REST API that contains
// Document struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type DocumentList struct {
	List
	Data []*Document `json:"data"`
}

// Find finds and returns Document with the given id. Returns nil if not found.
func (list *DocumentList) Find(id string) *Document {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}