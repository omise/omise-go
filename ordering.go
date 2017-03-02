package omise

// Ordering represents an enumeration of possible values for list ordering while
// performing operations against the Omise API that result in List objects being returned.
type Ordering string

// Ordering can be one of the following list of constants:
const (
	UnspecifiedOrder     Ordering = ""
	Chronological        Ordering = "chronological"
	ReverseChronological Ordering = "reverse_chronological"
)
