package omise

import (
	"time"
)

type PaymentReference struct {
	Base
	Barcode string `json:"barcode"`
	CustomerAmount int `json:"customer_amount"`
	CustomerCurrency string `json:"customer_currency"`
	CustomerExchangeRate int `json:"customer_exchange_rate"`
	DeviceID string `json:"device_id"`
	ExpiresAt time.Time `json:"expires_at"`
	OmiseTaxID string `json:"omise_tax_id"`
	PaymentCode string `json:"payment_code"`
	ReferenceNumber1 string `json:"reference_number_1"`
	ReferenceNumber2 string `json:"reference_number_2"`
	VaCode string `json:"va_code"`
}