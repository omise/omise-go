package omise

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"github.com/omise/omise-go/internal/creds"
)

var transport *http.Transport

func init() {
	certbytes, err := creds.Asset("ca_certificates.pem")
	if err != nil {
		panic(err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(certbytes) {
		panic("failed to load omise.co domain certificates.")
	}

	transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:    pool,
			MinVersion: tls.VersionTLS12,
		},
	}
}
