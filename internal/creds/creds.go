package creds

//go:generate go-bindata -nometadata -o ca_certificates.go -pkg creds ca_certificates.pem
//go:generate go fmt ca_certificates.go
