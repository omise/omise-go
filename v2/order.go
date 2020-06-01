package omise

// Order represents an enumeration of possible values for Search.
type Order string

// Order can be one of the following list of constants:
const (
	SearchChronological Order = "chronological"
	SearchReverseChronological Order = "reverse_chronological"
)

