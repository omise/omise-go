package operations

import (
	"github.com/omise/omise-go/internal"
	"net/url"
	"strconv"
	"time"
)

type CreateToken struct {
	Name            string
	Number          string
	ExpirationMonth time.Month
	ExpirationYear  int

	SecurityCode string
	City         string
	PostalCode   string
}

func (*CreateToken) Endpoint() (internal.Endpoint, string, string) {
	return internal.Vault, "POST", "/tokens"
}

func (c *CreateToken) Payload() (interface{}, error) {
	v := url.Values{}
	v.Set("card[name]", c.Name)
	v.Set("card[number]", c.Number)
	v.Set("card[expiration_month]", strconv.Itoa(int(c.ExpirationMonth)))
	v.Set("card[expiration_year]", strconv.Itoa(c.ExpirationYear))
	v.Set("card[security_code]", c.SecurityCode)
	v.Set("card[city]", c.City)
	v.Set("card[postal_code]", c.PostalCode)
	return v, nil
}
