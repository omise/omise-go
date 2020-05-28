package omise

// TransferList represents the list structure returned by Omise's REST API that contains
// Transfer struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type TransferList struct {
	Base
	Data []*Transfer `json:"data"`
}

// Find finds and returns Transfer with the given id. Returns nil if not found.
func (list *TransferList) Find(id string) *Transfer {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}