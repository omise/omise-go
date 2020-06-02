package omise

// OccurrenceList represents the list structure returned by Omise's REST API that contains
// Occurrence struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type OccurrenceList struct {
	List
	Data []*Occurrence `json:"data"`
}

// Find finds and returns Occurrence with the given id. Returns nil if not found.
func (list *OccurrenceList) Find(id string) *Occurrence {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}