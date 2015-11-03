package omise

import (
	"bytes"
	"encoding/json"
	"github.com/omise/omise-go/internal"
	"github.com/omise/omise-go/operations"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	*http.Client
	PublicKey string
	SecretKey string
}

func NewClient(pkey, skey string) (*Client, error) {
	switch {
	case pkey == "" && skey == "":
		return nil, ErrInvalidKey
	case pkey != "" && !strings.HasPrefix(pkey, "pkey_"):
		return nil, ErrInvalidKey
	case skey != "" && !strings.HasPrefix(skey, "skey_"):
		return nil, ErrInvalidKey
	}

	return &Client{&http.Client{}, pkey, skey}, nil
}

func (c *Client) Do(result interface{}, op internal.Operation) error {
	endpoint, method, path := op.Endpoint()

	// build request
	var body io.Reader
	payload, e := op.Payload()
	if e != nil {
		return e
	}

	if payload != nil {
		switch p := payload.(type) {
		case io.Reader:
			body = p
		case url.Values:
			body = ioutil.NopCloser(bytes.NewBuffer([]byte(p.Encode())))
		default:
			return ErrInternal("unsupported payload type.")
		}
	}

	req, e := http.NewRequest(method, string(endpoint)+path, body)
	if e != nil {
		return e
	}

	// authentication
	switch endpoint {
	case internal.API:
		req.SetBasicAuth(c.SecretKey, "")
	case internal.Vault:
		req.SetBasicAuth(c.PublicKey, "")
	default:
		return ErrInternal("unrecognized endpoint:" + endpoint)
	}

	// read response
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
		// postcond: status == 200 && e == nil
	}

	if result != nil {
		if json.NewDecoder(resp.Body).Decode(result); e != nil {
			return e
		}
	}

	return nil
}

// TODO: Generate this from list of types in the ./operations package.
func (c *Client) CreateCharge(op *operations.CreateCharge) (*Charge, error) {
	result := &Charge{}
	if e := c.Do(result, op); e != nil {
		return nil, e
	}

	return result, nil
}
