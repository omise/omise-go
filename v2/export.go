package omise

// Export represents resource object.
type Export struct {
	Base
	DownloadURI string `json:"download_uri"`
	FileType string `json:"file_type"`
	FilterParams map[string]interface{} `json:"filter_params"`
	FilterType string `json:"filter_type"`
	Location string `json:"location"`
	Name string `json:"name"`
	ObjectType string `json:"object_type"`
	Rows int64 `json:"rows"`
	Status string `json:"status"`
	Team string `json:"team"`
}

