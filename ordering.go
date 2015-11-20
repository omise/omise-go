package omise

// Ordering represents an enumeration of possible values for list ordering while
// performing operations against the Omise API that result in List objects being returned.
// Ordering can be one of the following list of constants:
type Ordering string

const (
	UnspecifiedOrder     Ordering = ""
	Chronological        Ordering = "chronological"
	ReverseChronological Ordering = "reverse_chronological"
)
