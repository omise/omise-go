package omise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

	// Overrides
	Endpoints map[internal.Endpoint]string

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

		Endpoints: map[internal.Endpoint]string{},
	}

	if len(build.Default.ReleaseTags) > 0 {
		client.GoVersion = build.Default.ReleaseTags[len(build.Default.ReleaseTags)-1]
	}

	return client, nil
}

// Request creates a new *http.Request that should performs the supplied Operation. Most
// people should use the Do method instead.
func (c *Client) Request(operation internal.Operation) (*http.Request, error) {
	var req *http.Request
	var err error
	if _, ok := operation.(json.Marshaler); ok {
		req, err = c.buildJSONRequest(operation)
	} else {
		req, err = c.buildFormRequest(operation)
	}

	if err != nil {
		return nil, err
	}

	err = c.setRequestHeaders(req, operation.Describe())
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) buildQuery(operation internal.Operation) (url.Values, error) {
	desc := operation.Describe()

	query, err := internal.MapURLValues(operation)
	if err != nil {
		return nil, err
	}

	if len(desc.Values) > 0 {
		for k, values := range desc.Values {
			if len(values) > 0 {
				query.Set(k, values[0])
			}
		}
	}

	return query, nil
}

func (c *Client) buildJSONRequest(operation internal.Operation) (*http.Request, error) {
	desc := operation.Describe()

	b, err := json.Marshal(operation)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(b)

	endpoint := string(desc.Endpoint)
	if ep, ok := c.Endpoints[desc.Endpoint]; ok {
		endpoint = ep
	}

	return http.NewRequest(desc.Method, endpoint+desc.Path, body)
}

func (c *Client) buildFormRequest(operation internal.Operation) (*http.Request, error) {
	desc := operation.Describe()

	query, err := c.buildQuery(operation)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if desc.Method != "GET" && desc.Method != "HEAD" {
		body = strings.NewReader(query.Encode())
	}

	endpoint := string(desc.Endpoint)
	if ep, ok := c.Endpoints[desc.Endpoint]; ok {
		endpoint = ep
	}

	req, err := http.NewRequest(desc.Method, endpoint+desc.Path, body)
	if err != nil {
		return nil, err
	}

	if desc.Method == "GET" || desc.Method == "HEAD" {
		req.URL.RawQuery = query.Encode()
	}

	return req, nil
}

func (c *Client) setRequestHeaders(req *http.Request, desc *internal.Description) error {
	ua := "OmiseGo/2015-11-06"
	if c.GoVersion != "" {
		ua += " Go/" + c.GoVersion
	}

	// Fallback between migrate to application/json
	if desc.ContentType == "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Add("Content-Type", desc.ContentType)
	}

	req.Header.Add("User-Agent", ua)
	if c.APIVersion != "" {
		req.Header.Add("Omise-Version", c.APIVersion)
	}

	switch desc.KeyKind() {
	case "public":
		req.SetBasicAuth(c.pkey, "")
	case "secret":
		req.SetBasicAuth(c.skey, "")
	default:
		return ErrInternal("unrecognized endpoint:" + desc.Endpoint)
	}

	return nil
}

// Do performs the supplied operation against Omise's REST API and unmarshal the response
// into the given result parameter. Results are usually basic objects or a list that
// corresponds to the operations being done.
//
// If the operation is successful, result should contains the response data. Otherwise a
// non-nil error should be returned. Error maybe of the omise-go.Error struct type, in
// which case you can further inspect the Code and Message field for more information.
func (c *Client) Do(result interface{}, operation internal.Operation) error {
	req, err := c.Request(operation)
	if err != nil {
		return err
	}

	// response
	resp, err := c.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &ErrTransport{err, buffer}
	}

	switch {
	case resp.StatusCode != 200:
		err := &Error{StatusCode: resp.StatusCode}
		if err := json.Unmarshal(buffer, err); err != nil {
			return &ErrTransport{err, buffer}
		}

		return err
	} // status == 200 && e == nil

	if c.debug {
		fmt.Println("resp:", resp.StatusCode, string(buffer))
	}

	if result != nil {
		if err := json.Unmarshal(buffer, result); err != nil {
			return &ErrTransport{err, buffer}
		}
	}

	return nil
}
