package operations_test

import (
	"testing"

	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	"github.com/omise/omise-go/operations"
	. "github.com/omise/omise-go/operations"
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

	exampleSource, createSource := &omise.Source{}, &operations.CreateSource{
		Type:     "fpx",
		Amount:   2000,
		Currency: "myr",
		Email:    "example@omise.co",
		Bank:     "ocbc",
	}

	client.MustDo(exampleSource, createSource)
	r.Equal(t, SourceID, exampleSource.ID)
}
