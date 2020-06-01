package omise

// RecipientList represents the list structure returned by Omise's REST API that contains
// Recipient struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type RecipientList struct {
	List
	Data []*Recipient `json:"data"`
}

// Find finds and returns Recipient with the given id. Returns nil if not found.
func (list *RecipientList) Find(id string) *Recipient {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
