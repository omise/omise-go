package omise

// FlowType represents an enumeration of possible values for Source.
type FlowType string

// FlowType can be one of the following list of constants:
const (
	SourceOffline FlowType = "offline"
	SourceRedirect FlowType = "redirect"
)

