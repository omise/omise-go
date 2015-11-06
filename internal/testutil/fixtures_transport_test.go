package testutil

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	a "github.com/stretchr/testify/assert"
)

type fixturesHTTPTest struct {
	method string
	url    string

	statusCode int
	filename   string
}

var fixtureTests = []*fixturesHTTPTest{
	{"GET", "https://api.omise.co/account", 200, "api.omise.co/account-get.json"},
	{"POST", "https://api.omise.co/charges", 200, "api.omise.co/charges-post.json"},
	{"GET", "https://api.omise.co/customers/not_exist", 404, "api.omise.co/customers/404-get.json"},
	{"GET", "https://vault.omise.co/tokens/tokn_test_4yq8lbecl0q6dsjzxr5", 200, "vault.omise.co/tokens/tokn_test_4yq8lbecl0q6dsjzxr5-get.json"},
}

func (test *fixturesHTTPTest) Test(t *testing.T) {
	wd, e := os.Getwd()
	if !a.NoError(t, e) {
		return
	}

	baseDir := filepath.Join(wd, FixturesDir)
	fixtures, e := NewFixturesTransport(baseDir)
	if !a.NoError(t, e) {
		return
	}

	t.Log(test.method, test.url)
	req, e := http.NewRequest(test.method, test.url, nil)
	if !a.NoError(t, e) {
		return
	}

	client := &http.Client{Transport: fixtures}
	if client == nil {
		panic("client nil")
	}
	resp, e := client.Do(req)
	if !(a.NoError(t, e) && a.NotNil(t, resp)) {
		return
	}

	a.Equal(t, test.statusCode, resp.StatusCode)

	defer resp.Body.Close()
	respBytes, e := ioutil.ReadAll(resp.Body)
	if !a.NoError(t, e) {
		return
	}

	fileBytes, e := ioutil.ReadFile(filepath.Join(wd, FixturesDir, test.filename))
	if !a.NoError(t, e) {
		return
	}

	a.Equal(t, string(fileBytes), string(respBytes))
}

func TestFixturesTransport(t *testing.T) {
	for _, test := range fixtureTests {
		test.Test(t)
	}
}
