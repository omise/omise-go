package omise

type WebhookDelivery struct {
	Base
	Status int `json:"status"`
	Uri string `json:"uri"`
}