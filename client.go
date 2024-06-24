package omise

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go/build"
	"io"
	"net/http"
	"strings"

	"github.com/omise/omise-go/internal"
)

// Library version to include in user agent
var libraryVersion = "1.5.0"

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
	APIVersion    string
	GoVersion     string
	customHeaders map[string]string
	ctx           context.Context
	userAgent     string
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

// WithCustomHeaders lets us add headers to request. This should be called before calling Do() method
func (c *Client) WithCustomHeaders(headers map[string]string) {
	if headers != nil {
		c.customHeaders = headers
	}
}

// WithContext By setting context, http request will use `NewRequestWithContext` which support to include tracing on same trace ID.
func (c *Client) WithContext(ctx context.Context) {
	if ctx != nil {
		c.ctx = ctx
	}
}

// WithUserAgent feature allows us to append additional userAgent information to the original user agent.
func (c *Client) WithUserAgent(userAgent string) {
	c.userAgent = userAgent
}

// If set to true then library will print the response of all the endpoints.
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// Request creates a new *http.Request that should performs the supplied Operation. Most
// people should use the Do method instead.
func (c *Client) Request(operation internal.Operation) (req *http.Request, err error) {
	req, err = c.buildJSONRequest(operation)
	if err != nil {
		return nil, err
	}

	err = c.setRequestHeaders(req, operation.Describe())
	if err != nil {
		return nil, err
	}

	return req, nil
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

	req, err := http.NewRequest(desc.Method, endpoint+desc.Path, body)
	if c.ctx != nil {
		req = req.WithContext(c.ctx)
	}
	return req, err
}

func (c *Client) setRequestHeaders(req *http.Request, desc *internal.Description) error {
	ua := c.userAgent
	ua += " OmiseGo/" + libraryVersion
	if c.GoVersion != "" {
		ua += " Go/" + c.GoVersion
	}

	if desc.ContentType != "" {
		req.Header.Add("Content-Type", desc.ContentType)
	}

	req.Header.Add("User-Agent", strings.TrimSpace(ua))
	if c.APIVersion != "" {
		req.Header.Add("Omise-Version", c.APIVersion)
	}

	// setting custom headers passed from the caller
	for k, v := range c.customHeaders {
		req.Header.Set(k, v)
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

	buffer, err := io.ReadAll(resp.Body)
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
