package omise

// Document represents Omise's document object.
// See https://www.omise.co/documents-api for more information.
type Document struct {
	Base
	Deleted     bool   `json:"deleted"`
	Filename    string `json:"filename"`
	DownloadURI string `json:"download_uri"`
}
