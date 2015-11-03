package omise

type Customer struct {
	Base
	DefaultCard string    `json:"default_card"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	Cards       *CardList `json:"cards"`
}
