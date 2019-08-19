package omise

import (
	"crypto/tls"
	"net/http"
)

var transport *http.Transport

func init() {
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
}
