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

// Based on this gist
// REF: https://gist.github.com/tonyhb/5819315
func MapURLValues(i interface{}) (url.Values, error) {
	result := url.Values{}
	if e := mapURLValues(i, result); e != nil {
		return nil, e
	}

	return result, nil
}

func mapURLValues(i interface{}, target url.Values) error {
	ival := reflect.ValueOf(i)
	switch ival.Kind() {
	case reflect.Interface, reflect.Ptr:
		ival = ival.Elem()
	}

	ityp := ival.Type()
	for i := 0; i < ival.NumField(); i++ {
		fieldval, field := ival.Field(i), ityp.Field(i)
		tag := field.Tag.Get("query")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		} else if tag == "-" {
			continue
		}

		// ptr deref
		out := ""
		if fieldval.Kind() == reflect.Ptr {
			if fieldval.IsNil() {
				continue
			}

			fieldval = fieldval.Elem()
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
			switch {
			case field.Type == timeType:
				t := fieldval.Interface().(time.Time)
				if !t.IsZero() {
					out = fieldval.Interface().(time.Time).Format(time.RFC3339Nano)
				}

			case field.Anonymous:
				if e := mapURLValues(fieldval.Interface(), target); e != nil {
					return e
				}

			default:
				return &ErrMap{field, "only embedded structs are mapped."}
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
