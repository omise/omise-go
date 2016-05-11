package internal

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
)

func Parameterize(operation Operation) (values url.Values, body io.Reader, op *Op, e error) {
	op = operation.Op()
	if values, e = MapURLValues(operation); e != nil {
		return
	}

	if len(op.Values) > 0 {
		for k, v := range op.Values {
			values[k] = v
		}
	}

	if !op.Multipart {
		switch op.Method {
		case "GET", "HEAD":
			body = nil
		default:
			body = bytes.NewBuffer([]byte(values.Encode()))
			values = url.Values{}
		}

	} else { // multipart
		buffer := &bytes.Buffer{}
		writer := multipart.NewWriter(buffer)
		for k, v := range values {
			if len(v) > 0 {
				writer.WriteField(k, v[0])
			}
		}

		if e = writer.Close(); e != nil {
			return
		}

		body = bytes.NewReader(buffer.Bytes())
		values = url.Values{}
	}

	return
}
