package internal

type Operation interface {
	Endpoint() (Endpoint, string, string)
	Payload() (interface{}, error)
}
