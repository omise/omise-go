package omise

// EventList represents the list structure returned by Omise's REST API that contains
// Event struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type EventList struct {
	List
	Data []*Event `json:"data"`
}

// Find finds and returns Event with the given id. Returns nil if not found.
func (list *EventList) Find(id string) *Event {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}