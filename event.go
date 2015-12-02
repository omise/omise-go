package omise

import (
	"encoding/json"
)

// Event represents Omise's event object.
type Event struct {
	Base
	Key  string      `json:"key" pretty:""`
	Data interface{} `json:"data" pretty:""`
}

type eventShim struct {
	Base
	key  string    `json:"key"`
	Data *Deletion `json:"data"`
}

// UnmarshalJSON unmarshals the buffer into an internal shim structure first, in order to
// determine the right structure to use for the .Data field. Then will re-unmarshal the
// structure as normal.
func (ev *Event) UnmarshalJSON(buffer []byte) error {
	shim := &eventShim{}
	if e := json.Unmarshal(buffer, shim); e != nil {
		return e
	}

	// go through a proxy type to undefine UnmarshalJSON (stack overflow, otherwise)
	type Event_ Event
	proxy := Event_(*ev)

	if shim.Data.Deleted {
		proxy.Data = shim.Data // already *Deletion

	} else { // Pre-init the right structure to match the returned type. TODO: Generate this?
		switch shim.Data.Object {
		case "charge":
			proxy.Data = &Charge{}
		case "customer":
			proxy.Data = &Customer{}
		case "card":
			proxy.Data = &Card{}
		case "dispute":
			proxy.Data = &Dispute{}
		case "recipient":
			proxy.Data = &Recipient{}
		case "refund":
			proxy.Data = &Refund{}
		case "transfer":
			proxy.Data = &Transfer{}
		}
	}

	if e := json.Unmarshal(buffer, &proxy); e != nil {
		return e
	}

	*ev = Event(proxy)
	return nil
}
