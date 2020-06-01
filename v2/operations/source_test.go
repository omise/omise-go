package operations_test

import (
	"testing"

	omise "github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

func TestSource(t *testing.T) {
	const (
		SourceID = "src_test_5a444nhh27tlyv81u40"
	)
	client := testutil.NewFixedClient(t)

	source := &omise.Source{}
	client.MustDo(source, &RetrieveSource{SourceID: SourceID})
	r.Equal(t, SourceID, source.ID)

	source = &omise.Source{}
	client.MustDo(source, &CreateSource{})
	r.Equal(t, SourceID, source.ID)
}
