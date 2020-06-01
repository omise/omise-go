package omise

// RecipientType represents an enumeration of possible values for Recipient.
type RecipientType string

// RecipientType can be one of the following list of constants:
const (
	RecipientCorporation RecipientType = "corporation"
	RecipientIndividual RecipientType = "individual"
)

