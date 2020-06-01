package operations_test

import (
	"testing"

	omise "github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

func TestSystemInfo(t *testing.T) {
	client := testutil.NewFixedClient(t)

	system := &omise.SystemInfo{}
	client.MustDo(system, &RetrieveSystemInfo{})

	r.Equal(t, "2019-05-29", system.Versions[0])
}
