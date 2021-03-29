package omise

// Customer represents Omise's customer object.
// See https://www.omise.co/customers-api for more information.
type Customer struct {
	Base
	DefaultCard string    `json:"default_card"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	Cards       *CardList `json:"cards"`

	Metadata map[string]interface{} `json:"metadata"`
}
