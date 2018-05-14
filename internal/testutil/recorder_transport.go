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
	_, fixpath, err := FixturePath(req)
	if err != nil {
		return nil, err
	}

	resp, err := transport.backing.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	file, err := os.Create(filepath.Join(FixtureBasePath, fixpath))
	if err != nil {
		return resp, err
	}

	reader := resp.Body
	defer reader.Close()

	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return resp, err
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(buffer))
	if _, err := io.Copy(file, bytes.NewBuffer(buffer)); err != nil {
		return resp, err
	}

	return resp, nil
}
