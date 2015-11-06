package testutil

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var _ = fmt.Println

type FixturesTransport struct {
	backing http.RoundTripper
	dir     string
}

func NewFixturesTransport(dir string) (*FixturesTransport, error) {
	if fi, e := os.Lstat(dir); e != nil {
		return nil, e
	} else if !fi.IsDir() {
		return nil, errors.New(dir + " is not a directory.")
	}

	backing := http.NewFileTransport(http.Dir(dir))
	return &FixturesTransport{backing, dir}, nil
}

func (transport *FixturesTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	is404, fixreq, fixurl := false, *req, *req.URL
	fixreq.URL = &fixurl

	urlpath := req.URL.Host + "/" +
		req.URL.Path[1:] + "-" + strings.ToLower(req.Method) + ".json"

	// construct the fixtures file path from the request
	filename := filepath.Join(transport.dir, urlpath)
	if _, e := os.Lstat(filename); e != nil {
		if !os.IsNotExist(e) {
			return nil, e
		}

		is404, urlpath = true, filepath.Join(
			filepath.Dir(urlpath),
			"404-"+strings.ToLower(req.Method)+".json",
		)
	}

	fixreq.URL.Path = urlpath

	// fetch the file
	resp, e := transport.backing.RoundTrip(&fixreq)
	if e != nil {
		return nil, e
	}

	// modify the response to looks legit.
	resp.Request = req
	if is404 {
		resp.StatusCode = 404
		resp.Status = "404 Not Found"
	}

	return resp, nil
}
