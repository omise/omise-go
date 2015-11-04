package omise

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/omise/omise-go/internal"
)

var _ = fmt.Println

type Client struct {
	*http.Client
	PublicKey string
	SecretKey string

	debug bool
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

	return &Client{&http.Client{}, pkey, skey, false}, nil
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
	if op.Method != "GET" && op.Method != "HEAD" {
		body = strings.NewReader(query.Encode())
	}

	if c.debug {
		fmt.Println(" req:", op.Method, string(op.Endpoint)+op.Path)
	}

	req, e := http.NewRequest(op.Method, string(op.Endpoint)+op.Path, body)
	// req, e := http.NewRequest(op.Method, "http://0.0.0.0:9999"+op.Path, body)
	if e != nil {
		return e
	}

	if op.Method == "GET" || op.Method == "HEAD" {
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
