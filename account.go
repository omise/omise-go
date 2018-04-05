package omise

// Account represents Omise's account object.
// See https://www.omise.co/account-api for more information.
type Account struct {
	Base
	Email string `json:"email"`
}
