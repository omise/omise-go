package omise

// ChainList represents the list structure returned by Omise's REST API that contains
// Chain struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type ChainList struct {
	List
	Data []*Chain `json:"data"`
}

// Find finds and returns Chain with the given id. Returns nil if not found.
func (list *ChainList) Find(id string) *Chain {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
