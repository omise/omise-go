package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	const Query = "amount:1000"
	client := testutil.NewFixedClient(t)

	result := &omise.ChargeSearchResult{}
	client.MustDo(result, &Search{
		Scope: omise.ChargeScope,
		Query: Query,
	})

	r.Equal(t, Query, result.Query)
	r.Equal(t, 1, result.Page)
	r.Equal(t, 3, result.TotalPages)

	charge := result.Data[0]
	r.Equal(t, "chrg_test_54i01932u4ts67cop81", charge.ID)
}

func TestSearch_Network(t *testing.T) {
	const ChargeID = "chrg_test_54i01932u4ts67cop81"

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// using query
	result := &omise.ChargeSearchResult{}
	client.MustDo(result, &Search{
		Scope: omise.ChargeScope,
		Query: ChargeID,
	})

	r.Equal(t, 1, result.Total)
	r.Equal(t, 1, result.TotalPages)

	charge := result.Data[0]
	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, int64(100000), charge.Amount)

	// using filters
	result = &omise.ChargeSearchResult{}
	client.MustDo(result, &Search{
		Scope: omise.ChargeScope,
		Query: "id:" + ChargeID,
	})

	r.Equal(t, 1, result.Total)
	r.Equal(t, 1, result.TotalPages)

	charge = result.Data[0]
	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, int64(100000), charge.Amount)
}
