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

// DateString represents date in string format used for payload e.g. "2020-01-01"
type DateString string

func (d DateString) String() string { return string(d) }

// MarshalJSON DateString type
func (d DateString) MarshalJSON() ([]byte, error) {
	s := string(d)
	if s == "" {
		return nil, nil
	}

	date, err := time.Parse("2006-01-02", string(s))
	if err != nil {
		return nil, err
	}

	return []byte(date.Format("\"2006-01-02\"")), nil
}
