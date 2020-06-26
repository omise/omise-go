package operations_test

import (
	"context"
	"testing"

	omise "github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/require"
)

func TestSystemInfo(t *testing.T) {
	client := testutil.NewFixedClient(t)

	system, _ := client.SystemInfo().Retrieve(context.Background(), &omise.RetrieveSystemInfoParams{})

	r.Equal(t, "2019-05-29", system.Versions[0])
}
