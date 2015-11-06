package omise

import (
	"encoding/json"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/omise/omise-go/internal"
)

var _ = fmt.Println

// Client helps you configure and perform HTTP operations against Omise's REST API. It
// should be used with operation structures from the operations subpackage.
type Client struct {
	*http.Client
	debug bool
	pkey  string
	skey  string

	// configuration
	APIVersion string
	GoVersion  string
}

// NewClient creates and returns a Client with the given public key and secret key.  Signs
// in to http://omise.co and visit https://dashboard.omise.co/test/dashboard to obtain
// your test (or live) keys.
func NewClient(pkey, skey string) (*Client, error) {
	switch {
	case pkey == "" && skey == "":
		return nil, ErrInvalidKey
	case pkey != "" && !strings.HasPrefix(pkey, "pkey_"):
		return nil, ErrInvalidKey
	case skey != "" && !strings.HasPrefix(skey, "skey_"):
		return nil, ErrInvalidKey
	}

	client := &Client{
		Client: &http.Client{Transport: transport},
		debug:  false,
		pkey:   pkey,
		skey:   skey,
	}

	if len(build.Default.ReleaseTags) > 0 {
		client.GoVersion = build.Default.ReleaseTags[len(build.Default.ReleaseTags)-1]
	}

	return client, nil
}

// Request creates a new *http.Request that should performs the supplied Operation. Most
// people should use the Do method instead.
func (c *Client) Request(operation internal.Operation) (*http.Request, error) {
	op := operation.Op()

	query, e := internal.MapURLValues(operation)
	if e != nil {
		return nil, e
	}

	if len(op.Values) > 0 {
		for k, values := range op.Values {
			if len(values) > 0 {
				query.Set(k, values[0])
			}
		}
	}

	body := io.Reader(nil)
	if op.Method != "GET" && op.Method != "HEAD" {
		body = strings.NewReader(query.Encode())
	}

	if c.debug {
		fmt.Println(" req:", op.Method, string(op.Endpoint)+op.Path)
	}

	req, e := http.NewRequest(op.Method, string(op.Endpoint)+op.Path, body)
	// req, e := http.NewRequest(op.Method, "http://0.0.0.0:9999"+op.Path, body)
	if e != nil {
		return nil, e
	}

	if op.Method == "GET" || op.Method == "HEAD" {
		req.URL.RawQuery = query.Encode()
	}

	// headers
	ua := "OmiseGo/2015-11-06"
	if c.GoVersion != "" {
		ua += " Go/" + c.GoVersion
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", ua)
	if c.APIVersion != "" {
		req.Header.Add("Omise-Version", c.APIVersion)
	}

	switch op.Endpoint {
	case internal.API:
		req.SetBasicAuth(c.skey, "")
	case internal.Vault:
		req.SetBasicAuth(c.pkey, "")
	default:
		return nil, ErrInternal("unrecognized endpoint:" + op.Endpoint)
	}

	return req, nil
}

// Do performs the supplied operation against Omise's REST API and unmarshal the response
// into the given result parameter. Results are usually basic objects or a list that
// corresponds to the operations being done.
//
// If the operation is successful, result should contains the response data. Otherwise a
// non-nil error should be returned. Error maybe of the omise-go.Error struct type, in
// which case you can further inspect the Code and Message field for more information.
func (c *Client) Do(result interface{}, operation internal.Operation) error {
	req, e := c.Request(operation)
	if e != nil {
		return e
	}

	// response
	resp, e := c.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	switch {
	case e != nil:
		return e
	case resp.StatusCode != 200:
		err := &Error{StatusCode: resp.StatusCode}
		if e := json.NewDecoder(resp.Body).Decode(err); e != nil {
			return e // TODO: Wrap so body is never lost.
		}
		return err
	} // status == 200 && e == nil

	buffer, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}

	if c.debug {
		fmt.Println("resp:", resp.StatusCode, string(buffer))
	}

	if result != nil {
		if e := json.Unmarshal(buffer, result); e != nil {
			return e
		}
	}

	return nil
}
