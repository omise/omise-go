package omise

import (
	"encoding/json"
	"time"
)

// Date type for marshal/unmarshal JSON without time
type Date time.Time

func (d Date) String() string { return time.Time(d).Format("2006-01-02") }

// UnmarshalJSON Date type
func (d *Date) UnmarshalJSON(b []byte) error {
	tm, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return json.Unmarshal(b, (*time.Time)(d))
	}
	*d = Date(tm)
	return nil
}

// MarshalJSON Date type
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format("\"2006-01-02\"")), nil
}
