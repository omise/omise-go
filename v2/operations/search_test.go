package operations_test

import (
	"context"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	const Query = "amount:1000"
	client := testutil.NewFixedClient(t)

	result, _ := client.Search().SearchCharges(context.Background(), &omise.SearchParams{
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
	result, _ := client.Search().SearchCharges(context.Background(), &omise.SearchParams{
		Query: ChargeID,
	})

	r.Equal(t, 1, result.Total)
	r.Equal(t, 1, result.TotalPages)

	charge := result.Data[0]
	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, int64(100000), charge.Amount)

	// using filters
	result, _ = client.Search().SearchCharges(context.Background(), &omise.SearchParams{
		Query: "id:" + ChargeID,
	})

	r.Equal(t, 1, result.Total)
	r.Equal(t, 1, result.TotalPages)

	charge = result.Data[0]
	r.Equal(t, ChargeID, charge.ID)
	r.Equal(t, int64(100000), charge.Amount)
}
