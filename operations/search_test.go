package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	const Query = "amount:1000"
	client := testutil.NewFixedClient(t)

	result := &omise.ChargeSearchResult{}
	client.MustDo(result, &Search{
		Scope: omise.ChargeScope,
		Query: Query,
	})

	a.Equal(t, Query, result.Query)
	a.Equal(t, 1, result.Page)
	a.Equal(t, 3, result.TotalPages)

	charge := result.Data[0]
	a.Equal(t, "chrg_test_54i01932u4ts67cop81", charge.ID)
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

	a.Equal(t, 1, result.Total)
	a.Equal(t, 1, result.TotalPages)

	charge := result.Data[0]
	a.Equal(t, ChargeID, charge.ID)
	a.Equal(t, int64(100000), charge.Amount)

	// using filters
	result = &omise.ChargeSearchResult{}
	client.MustDo(result, &Search{
		Scope: omise.ChargeScope,
		Query: "id:" + ChargeID,
	})

	a.Equal(t, 1, result.Total)
	a.Equal(t, 1, result.TotalPages)

	charge = result.Data[0]
	a.Equal(t, ChargeID, charge.ID)
	a.Equal(t, int64(100000), charge.Amount)
}
