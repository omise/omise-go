package operations_test

import (
	"context"
	"testing"
	"time"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

func TestLink_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// create
	link, _ := client.Link().Create(context.Background(), &omise.CreateLinkParams{
		Amount:      99900,
		Currency:    "thb",
		Title:       "Hot Latte",
		Description: "A warm cup of coffee",
		Multiple:    true,
	})
	r.NotNil(t, link)

	t.Log("created link:", link.ID)
	r.Equal(t, int64(99900), link.Amount)
	r.Equal(t, "Hot Latte", link.Title)
	r.True(t, link.Multiple)

	// retrieve created link
	link2, _ := client.Link().Retrieve(context.Background(), &omise.RetrieveLinkParams{LinkID: link.ID})
	r.Equal(t, link.Amount, link2.Amount)
	r.Equal(t, link.Title, link2.Title)

	// list created links from the last hour
	links, list := &omise.LinkList{}, &ListLinks{
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	client.MustDo(links, list)

	r.True(t, len(links.Data) > 0, "links list empty!")
	link2 = links.Find(link.ID)
	r.NotNil(t, link2, "could not find recent links in list.")
	r.Equal(t, link.ID, link2.ID)
	r.Equal(t, link.Amount, link2.Amount)
	r.Equal(t, link.Title, link2.Title)
}
