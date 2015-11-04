package omise

type Customer struct {
	Base
	DefaultCard string    `json:"default_card"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	Cards       *CardList `json:"cards"`
}

type CustomerList struct {
	List
	Data []*Customer `json:"data"`
}

func (list *CustomerList) Find(id string) *Customer {
	for _, customer := range list.Data {
		if customer.ID == id {
			return customer
		}
	}

	return nil
}
