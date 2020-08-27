package omise

// TransferScheduleList represents the list structure returned by Omise's REST API that contains
// TransferSchedule struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type TransferScheduleList struct {
	List
	Data []*TransferSchedule `json:"data"`
}

// Find finds and returns TransferSchedule with the given id. Returns nil if not found.
func (list *TransferScheduleList) Find(id string) *TransferSchedule {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
