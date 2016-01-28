package testutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type RecorderTransport struct {
	backing http.RoundTripper
}

func NewRecorderTransport() (*RecorderTransport, error) {
	return &RecorderTransport{http.DefaultTransport}, nil
}

func (transport *RecorderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// TODO: Allow recording to new non-existent file.
	_, fixpath, e := FixturePath(req)
	if e != nil {
		return nil, e
	}

	resp, e := transport.backing.RoundTrip(req)
	if e != nil {
		return resp, e
	}

	file, e := os.Create(filepath.Join(FixtureBasePath, fixpath))
	if e != nil {
		return resp, e
	}

	reader := resp.Body
	defer reader.Close()

	buffer, e := ioutil.ReadAll(reader)
	if e != nil {
		return resp, e
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(buffer))
	if _, e := io.Copy(file, bytes.NewBuffer(buffer)); e != nil {
		return resp, e
	}

	return resp, nil
}
