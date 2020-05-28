package omise

type SystemInfo struct {
	Base
	Location string `json:"location"`
	Versions []string `json:"versions"`
}
