package omise

// ScheduleList represents the list structure returned by Omise's REST API that contains
// Schedule struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type ScheduleList struct {
	Base
	Data []*Schedule `json:"data"`
}

// Find finds and returns Schedule with the given id. Returns nil if not found.
func (list *ScheduleList) Find(id string) *Schedule {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}