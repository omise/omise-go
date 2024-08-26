package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example: Create source
//
//	source, createSource := &omise.Source{}, &CreateSource{
//		Amount:   100000,
//		Currency: "thb",
//		Type:     "bill_payment_tesco_lotus",
//	}
//	if e := client.Do(source, createSource); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("created source:", source.ID)
type CreateSource struct {
	Type                     string          `json:"type"`
	Amount                   int64           `json:"amount"`
	Currency                 string          `json:"currency"`
	Bank                     string          `json:"bank,omitempty"`
	Barcode                  string          `json:"barcode,omitempty"`
	Email                    string          `json:"email,omitempty"`
	InstallmentTerm          int64           `json:"installment_term,omitempty"`
	Name                     string          `json:"name,omitempty"`
	MobileNumber             string          `json:"mobile_number,omitempty"`
	StoreID                  string          `json:"store_id,omitempty"`
	StoreName                string          `json:"store_name,omitempty"`
	TerminalID               string          `json:"terminal_id,omitempty"`
	ZeroInterestInstallments bool            `json:"zero_interest_installments,omitempty"`
	PlatformType             string          `json:"platform_type,omitempty"`
	Ip                       string          `json:"ip,omitempty"`
	Billing                  BillingShipping `json:"-"`
	Shipping                 BillingShipping `json:"-"`
	PromotionCode            string          `json:"promotion_code,omitempty"`
	Items                    Items           `json:"-"`
}

type BillingShipping struct {
	Country    string `json:"country"`
	City       string `json:"City"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Street1    string `json:"street1"`
	Street2    string `json:"street2,omitempty"`
}

type Items struct {
	Amount   int64  `json:"amount"`
	Sku      string `json:"sku,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	Category string `json:"category,omitempty"`
	Brand    string `json:"brand,omitempty"`
	ItemUri  string `json:"item_uri,omitempty"`
	ImageUri string `json:"image_uri,omitempty"`
}

func (req *CreateSource) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/sources",
		ContentType: "application/json",
		APIKey:      "public",
	}
}

// Example: Retrieve Source
//
//	source, retrieve := &omise.Source{}, &RetrieveSource{"src_123"}
//	if e := client.Do(source, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("source #123: %#v\n", source)
type RetrieveSource struct {
	SourceID string `json:"-"`
}

func (req *RetrieveSource) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/sources/" + req.SourceID,
		ContentType: "application/json",
	}
}
