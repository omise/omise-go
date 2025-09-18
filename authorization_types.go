package omise

// AuthorizationType represents an enumeration of a possible authorization type of a Charge object.
type AuthorizationType string

// AuthorizationType can be one of the following list of constants:
const (
	PreAuth   AuthorizationType = "pre_auth"
	FinalAuth AuthorizationType = "final_auth"
)
