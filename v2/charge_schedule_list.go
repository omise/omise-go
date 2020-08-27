package omise

// ChargeScheduleList represents the list structure returned by Omise's REST API that contains
// ChargeSchedule struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type ChargeScheduleList struct {
	List
	Data []*ChargeSchedule `json:"data"`
}

// Find finds and returns ChargeSchedule with the given id. Returns nil if not found.
func (list *ChargeScheduleList) Find(id string) *ChargeSchedule {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
