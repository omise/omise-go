package internal

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var timeType = reflect.TypeOf(time.Time{})

type ErrMap struct {
	field  reflect.StructField
	reason string
}

func (e *ErrMap) Error() string {
	return "cannot map field `" + e.field.Name + "`, " + e.reason
}

// MapURLValues maps a user-defined struct to url.Values. Nested structs, arrays, and maps
// are mapped to field name with "[]" suffix with the current index as map key.
func MapURLValues(i interface{}) (url.Values, error) {
	result := url.Values{}
	if e := mapURLValues(i, result, ""); e != nil {
		return nil, e
	}

	return result, nil
}

func mapURLValues(i interface{}, target url.Values, parent string) error {
	ival := reflect.ValueOf(i)
	switch ival.Kind() {
	case reflect.Interface, reflect.Ptr:
		ival = ival.Elem()
	}

	ityp := ival.Type()
	for i := 0; i < ival.NumField(); i++ {
		fieldval, field := ival.Field(i), ityp.Field(i)
		tag, sendZero := "", false

		// compute tag names and options
		opt, opts := "", strings.Split(field.Tag.Get("query"), ",")
		if len(opts) > 0 {
			tag, opts = opts[0], opts[1:]
		}
		if tag == "" {
			tag = field.Tag.Get("json")
		}
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}

		if tag == "-" {
			continue
		}
		if parent != "" {
			tag = parent + "[" + tag + "]"
		}

		for len(opts) > 0 {
			opt, opts = opts[0], opts[1:]
			switch opt {
			case "sendzero":
				sendZero = true
			}
		}

		// ptr deref
		out := ""
		if fieldval.Kind() == reflect.Ptr {
			if fieldval.IsNil() {
				continue
			}

			fieldval = fieldval.Elem()
		}

		// zero check
		isZero := fieldval.Interface() == reflect.Zero(fieldval.Type()).Interface()
		if isZero && !sendZero {
			continue
		}

		// convert
		switch fieldval.Kind() {
		case reflect.Bool:
			out = strconv.FormatBool(fieldval.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			out = strconv.FormatInt(fieldval.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			out = strconv.FormatUint(fieldval.Uint(), 10)
		case reflect.Float32:
			out = strconv.FormatFloat(fieldval.Float(), 'f', 4, 32)
		case reflect.Float64:
			out = strconv.FormatFloat(fieldval.Float(), 'f', 4, 64)
		case reflect.String:
			out = fieldval.String()

		case reflect.Struct:
			switch { // well-known types
			case field.Type == timeType:
				t := fieldval.Interface().(time.Time)
				if !t.IsZero() {
					out = fieldval.Interface().(time.Time).Format(time.RFC3339Nano)
				}

			case field.Anonymous: // embedded structs
				if e := mapURLValues(fieldval.Interface(), target, ""); e != nil {
					return e
				}

			default: // named struct fields
				if e := mapURLValues(fieldval.Interface(), target, tag); e != nil {
					return e
				}
			}

		default:
			// TODO: More types.
			return &ErrMap{field, "unsupported type."}
		}

		if out != "" {
			target.Set(tag, out)
		}
	}

	return nil
}
