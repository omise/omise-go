package omise

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/omise/omise-go/internal"
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

func (c *Client) Do(result interface{}, operation internal.Operation) error {
	op := operation.Op()

	// request
	query := internal.MapURLValues(operation)
	if len(op.Values) > 0 {
		for k, values := range op.Values {
			if len(values) > 0 {
				query.Set(k, values[0])
			}
		}
	}

	body := io.Reader(nil)
	if len(query) > 0 {
		body = ioutil.NopCloser(strings.NewReader(query.Encode()))
	}

	req, e := http.NewRequest(op.Method, string(op.Endpoint)+op.Path, body)
	if e != nil {
		return e
	}

	switch op.Endpoint {
	case internal.API:
		req.SetBasicAuth(c.SecretKey, "")
	case internal.Vault:
		req.SetBasicAuth(c.PublicKey, "")
	default:
		return ErrInternal("unrecognized endpoint:" + op.Endpoint)
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

	if result != nil {
		if json.NewDecoder(resp.Body).Decode(result); e != nil {
			return e
		}
	}

	return nil
}
