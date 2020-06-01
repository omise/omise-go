package testutil

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	r "github.com/stretchr/testify/require"
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
	fixtures, err := NewFixturesTransport()
	r.NoError(t, err)

	t.Log(test.method, test.url)
	req, err := http.NewRequest(test.method, test.url, nil)
	r.NoError(t, err)

	client := &http.Client{Transport: fixtures}
	resp, err := client.Do(req)
	r.NoError(t, err)
	r.NotNil(t, resp)
	r.Equal(t, test.statusCode, resp.StatusCode)

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	r.NoError(t, err)

	fixBytes, err := ioutil.ReadFile(filepath.Join(FixtureBasePath, test.filename))
	r.NoError(t, err)
	r.Equal(t, string(fixBytes), string(respBytes))
}

func TestFixturesTransport(t *testing.T) {
	for _, test := range fixtureTests {
		test.Test(t)
	}
}
