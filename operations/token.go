package operations

import (
	"time"

	"github.com/omise/omise-go/internal"
)

type CreateToken struct {
	Name            string     `query:"card[name]"`
	Number          string     `query:"card[number]"`
	ExpirationMonth time.Month `query:"card[expiration_month]"`
	ExpirationYear  int        `query:"card[expiration_year]"`

	SecurityCode string `query:"card[security_code]"`
	City         string `query:"card[city]"`
	PostalCode   string `query:"card[postal_code]"`
}

func (req *CreateToken) Op() *internal.Op {
	return &internal.Op{internal.Vault, "POST", "/tokens", nil}
}

type RetreiveToken struct {
	ID string `query:"-"`
}

func (token *RetreiveToken) Op() *internal.Op {
	return &internal.Op{internal.Vault, "GET", "/tokens/" + token.ID, nil}
}
