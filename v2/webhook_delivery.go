package omise

// WebhookDelivery represents resource object.
type WebhookDelivery struct {
	Base
	Status int `json:"status"`
	URI string `json:"uri"`
}

