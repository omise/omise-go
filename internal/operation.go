package internal

type Operation interface {
	Endpoint() (Endpoint, string, string)
}

type PayloadOperation interface {
	Operation
	Payload() (interface{}, error)
}
