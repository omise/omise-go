package omise

// Order represents an enumeration of possible values for List.
type Order string

// Order can be one of the following list of constants:
const (
	ListChronological Order = "chronological"
	ListReverseChronological Order = "reverse_chronological"
)

