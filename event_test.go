package omise_test

import (
	"testing"

	. "github.com/omise/omise-go"
	r "github.com/stretchr/testify/require"
)

func TestEvent_UnmarshalJSON(t *testing.T) {
	type args struct {
		buffer []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should unmarshal event",
			args: args{
				buffer: []byte(`{"key":"charge.complete", "data":{"object":"charge", "deleted":true}}`),
			},
		},
		{
			name: "should unmarshal empty event",
			args: args{
				buffer: []byte(`{}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := &Event{}
			err := ev.UnmarshalJSON(tt.args.buffer)
			r.NoError(t, err)
		})
	}
}
