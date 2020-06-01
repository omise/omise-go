package omise

// Customer represents Customer object.
// See https://www.omise.co/customers-api for more information.
type Customer struct {
	Base
	Cards *CardList `json:"cards"`
	DefaultCard string `json:"default_card"`
	Description string `json:"description"`
	Email string `json:"email"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
}

