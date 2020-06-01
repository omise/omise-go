package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	account, retrieve := &omise.Account{}, &RetrieveAccount{
//	}
//	if e := client.Do(account, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("account: %#v\n", account)
//
type RetrieveAccount struct {
}

// Describe
func (req *RetrieveAccount) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/account",
		ContentType: "application/json",
	}
}

// Example:
//
//	account, update := &omise.Account{}, &UpdateAccount{
//	}
//	if e := client.Do(account, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("account: %#v\n", account)
//
type UpdateAccount struct {
	ChainEnabled bool `json:"chain_enabled,omitempty"`
	ChainReturnURI string `json:"chain_return_uri,omitempty"`
	MetadataExportKeys map[string]interface{} `json:"metadata_export_keys,omitempty"`
	WebhookURI string `json:"webhook_uri,omitempty"`
	ZeroInterestInstallments bool `json:"zero_interest_installments,omitempty"`
}

// MetadataExportKeys
type MetadataExportKeys struct {
	Charge []interface{} `json:"charge,omitempty"`
	Dispute []interface{} `json:"dispute,omitempty"`
	Refund []interface{} `json:"refund,omitempty"`
	Transfer []interface{} `json:"transfer,omitempty"`
}

// Describe
func (req *UpdateAccount) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/account",
		ContentType: "application/json",
	}
}

