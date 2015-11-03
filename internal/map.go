package internal

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Based on this gist
// REF: https://gist.github.com/tonyhb/5819315
func MapURLValues(i interface{}) url.Values {
	result := url.Values{}

	ival := reflect.ValueOf(i).Elem()
	typ := ival.Type()
	for i := 0; i < ival.NumField(); i++ {
		field, tag := ival.Field(i), typ.Field(i).Tag.Get("query")
		if tag == "" {
			tag = strings.ToLower(typ.Field(i).Name)
		} else if tag == "-" {
			continue
		}

		// ptr deref
		out := ""
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				continue
			}

			field = field.Elem()
		}

		// convert
		switch field.Kind() {
		case reflect.Bool:
			out = strconv.FormatBool(field.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			out = strconv.FormatInt(field.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			out = strconv.FormatUint(field.Uint(), 10)
		case reflect.Float32:
			out = strconv.FormatFloat(field.Float(), 'f', 4, 32)
		case reflect.Float64:
			out = strconv.FormatFloat(field.Float(), 'f', 4, 64)
		case reflect.String:
			out = field.String()
			// TODO: More types.
		}
		if out != "" {
			result.Set(tag, out)
		}
	}

	return result
}
