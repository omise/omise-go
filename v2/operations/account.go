package operations


// Example:
//
//	charge, update := &omise.Account{}, &RetrieveAccount{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveAccount struct {
	Base
}

func (req *RetrieveAccount) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/account",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Account{}, &UpdateAccount{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateAccount struct {
	Base
	ChainEnabled bool `json:"chain_enabled"`
	ChainReturnURI string `json:"chain_return_uri"`
	MetadataExportKeys *MetadataExportKeys `json:"metadata_export_keys"`
	WebhookURI string `json:"webhook_uri"`
	ZeroInterestInstallments bool `json:"zero_interest_installments"`
}

func (req *UpdateAccount) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      PATCH,
		Path:        "/account",
		ContentType: "application/json",
	}
}

