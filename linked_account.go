package omise

import "time"

// LinkedAccount is the representation of the linked_account object
type LinkedAccount struct {
	Base

	RegistrationURI string `json:"registration_uri"`
	ReturnURI       string `json:"return_uri"`

	Type       string    `json:"type"`
	Status     string    `json:"status"`
	LastDigits int       `json:"last_digits"`
	ExpiresAt  time.Time `json:"expires_at"`

	FailureCode    *string    `json:"failure_code"`
	FailureMessage *string    `json:"failure_message"`
	RegisteredAt   *time.Time `json:"registered_at,omitempty"`

	Metadata map[string]interface{} `json:"metadata"`
}
