package omise

// WebhookDelivery represents resource object.
type WebhookDelivery struct {
	Base
	Status int `json:"status"`
	URI string `json:"uri"`
}

// WebhookDeliveryService represents resource service.
type WebhookDeliveryService struct {
	*Client
}

// WebhookDelivery defines resource service.
func (c *Client) WebhookDelivery() *WebhookDeliveryService {
	return &WebhookDeliveryService{c}
}

