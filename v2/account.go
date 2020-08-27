package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Account represents Account object.
// See https://www.omise.co/account-api for more information.
type Account struct {
	Base
	APIVersion string `json:"api_version"`
	AutoActivateRecipients bool `json:"auto_activate_recipients"`
	ChainEnabled bool `json:"chain_enabled"`
	ChainReturnURI string `json:"chain_return_uri"`
	Country string `json:"country"`
	Currency string `json:"currency"`
	Email string `json:"email"`
	Location string `json:"location"`
	MetadataExportKeys map[string]interface{} `json:"metadata_export_keys"`
	SupportedCurrencies []string `json:"supported_currencies"`
	Team string `json:"team"`
	WebhookURI string `json:"webhook_uri"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}

// AccountService represents resource service.
type AccountService struct {
	*Client
}

// Account defines resource service.
func (c *Client) Account() *AccountService {
	return &AccountService{c}
}

// Retrieve retrieves account
//
// Example:
//
//	account, err := client.Account().Retrieve(ctx, &RetrieveAccountParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("account: %#v\n", account)
//
func (s *AccountService) Retrieve(ctx context.Context, params *RetrieveAccountParams) (*Account, error) {
	result := &Account{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveAccountParams params object.
type RetrieveAccountParams struct {
}

// Describe describes structure of request
func (req *RetrieveAccountParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/account",
		ContentType: "application/json",
	}
}

// Update updates account
//
// Example:
//
//	account, err := client.Account().Update(ctx, &UpdateAccountParams{
//	})
//	if err != nil {
//		panic(e)
//	}
//
//	fmt.Printf("account: %#v\n", account)
//
func (s *AccountService) Update(ctx context.Context, params *UpdateAccountParams) (*Account, error) {
	result := &Account{}
	err := s.Do(ctx, result, params)

	return result, err
}

// UpdateAccountParams params object.
type UpdateAccountParams struct {
	ChainEnabled bool `json:"chain_enabled,omitempty"`
	ChainReturnURI string `json:"chain_return_uri,omitempty"`
	MetadataExportKeys map[string]interface{} `json:"metadata_export_keys,omitempty"`
	WebhookURI string `json:"webhook_uri,omitempty"`
	ZeroInterestInstallments bool `json:"zero_interest_installments,omitempty"`
}

// MetadataExportKeysParams params object.
type MetadataExportKeysParams struct {
	Charge []interface{} `json:"charge,omitempty"`
	Dispute []interface{} `json:"dispute,omitempty"`
	Refund []interface{} `json:"refund,omitempty"`
	Transfer []interface{} `json:"transfer,omitempty"`
}

// Describe describes structure of request
func (req *UpdateAccountParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/account",
		ContentType: "application/json",
	}
}

