package omise

// Document represents Document object.
// See https://www.omise.co/documents-api for more information.
type Document struct {
	Base
	DownloadURI string `json:"download_uri"`
	Filename string `json:"filename"`
	Location string `json:"location"`
}