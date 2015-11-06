package omise

// RecipientType represents an enumeration of possible types of Recipient(s) which can be
// one of the following list of constants:
type RecipientType string

const (
	Individual  RecipientType = "individual"
	Corporation RecipientType = "corporation"
)
