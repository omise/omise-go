package testutil

import (
	"fmt"
	"net/http"
)

var _ = fmt.Println

type FixturesTransport struct {
	backing http.RoundTripper
}

func NewFixturesTransport() (*FixturesTransport, error) {
	backing := http.NewFileTransport(http.Dir(FixtureBasePath))
	return &FixturesTransport{backing}, nil
}

func (transport *FixturesTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fixreq, fixurl := *req, *req.URL
	fixreq.URL = &fixurl

	status, fixpath, e := FixturePath(req)
	if e != nil {
		return nil, e
	}

	// fetch the file
	fixreq.URL.Path = fixpath

	resp, e := transport.backing.RoundTrip(&fixreq)
	if e != nil {
		return nil, e
	}

	// modify the FS transport response to looks like it came remotely.
	resp.Request = req
	resp.StatusCode = status
	switch status {
	case 404:
		resp.Status = "404 Not Found"
	}

	return resp, nil
}
