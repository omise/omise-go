package omise

// Event represents Event object.
// See https://www.omise.co/events-api for more information.
type Event struct {
	Base
	Data map[string]interface{} `json:"data"`
	Key string `json:"key"`
	Location string `json:"location"`
	WebhookDeliveries []interface{} `json:"webhook_deliveries"`
}