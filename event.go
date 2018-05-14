package omise

import (
	"encoding/json"
)

// Event represents Omise's event object.
type Event struct {
	Base
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}

type eventShim struct {
	Base
	Key  string    `json:"key"`
	Data *Deletion `json:"data"`
}

// UnmarshalJSON unmarshals the buffer into an internal shim structure first, in order to
// determine the right structure to use for the .Data field. Then will re-unmarshal the
// structure as normal.
func (ev *Event) UnmarshalJSON(buffer []byte) error {
	shim := &eventShim{}
	if err := json.Unmarshal(buffer, shim); err != nil {
		return err
	}

	// go through a proxy type to undefine UnmarshalJSON (stack overflow, otherwise)
	type EventProxy Event
	proxy := EventProxy(*ev)
	proxy.Key = shim.Key

	// Pre-init the right structure to match the returned type.
	if shim.Data.Deleted {
		proxy.Data = shim.Data // already *Deletion
	} else {
		proxy.Data = ev.dataInstanceFromType(shim.Data.Object)
	}

	if err := json.Unmarshal(buffer, &proxy); err != nil {
		return err
	}

	*ev = Event(proxy)
	return nil
}

func (ev *Event) dataInstanceFromType(typ string) interface{} {
	// TODO: Generate this.
	switch typ {
	case "charge":
		return &Charge{}
	case "customer":
		return &Customer{}
	case "card":
		return &Card{}
	case "dispute":
		return &Dispute{}
	case "recipient":
		return &Recipient{}
	case "refund":
		return &Refund{}
	case "transfer":
		return &Transfer{}
	}

	return nil
}
