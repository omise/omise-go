package omise

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

