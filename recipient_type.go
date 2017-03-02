package omise

// RecipientType represents an enumeration of possible types of Recipient(s).
type RecipientType string

// RecipientType can be one of the following list of constants:
const (
	Individual  RecipientType = "individual"
	Corporation RecipientType = "corporation"
)
