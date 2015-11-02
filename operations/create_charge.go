package operations

import (
	"github.com/omise/omise-go/internal"
	"net/url"
	"strconv"
)

type CreateCharge struct {
	Customer    string
	Card        string
	Amount      int64
	Currency    string
	Description string
	DontCapture bool
	ReturnURI   string
}

func (*CreateCharge) Endpoint() (internal.Endpoint, string, string) {
	return internal.API, "POST", "/charges"
}

func (c *CreateCharge) Payload() (interface{}, error) {
	values := url.Values{}
	if c.Customer != "" {
		values.Set("customer", c.Customer)
	}
	if c.Card != "" {
		values.Set("card", c.Card)
	}

	values.Add("amount", strconv.FormatInt(c.Amount, 10))
	values.Add("currency", c.Currency)

	if c.Description != "" {
		values.Add("description", c.Description)
	}
	if c.DontCapture {
		values.Add("capture", "false")
	}
	if c.ReturnURI != "" {
		values.Add("return_uri", c.ReturnURI)
	}

	return values, nil
}
