package omise

// SystemInfo represents resource object.
type SystemInfo struct {
	Base
	Location string `json:"location"`
	Versions []string `json:"versions"`
}

