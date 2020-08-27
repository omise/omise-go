package omise

import (
	"encoding/json"
	"time"
)

// MarshalJSON List type
func (l ListParams) MarshalJSON() ([]byte, error) {
	type Alias ListParams
	params := struct {
		Alias
		PFrom *time.Time `json:"from,omitempty"`
		PTo   *time.Time `json:"to,omitempty"`
	}{
		Alias: Alias(l),
	}
	if !l.From.IsZero() {
		params.PFrom = &l.From
	}
	if !l.To.IsZero() {
		params.PTo = &l.To
	}
	return json.Marshal(params)
}
