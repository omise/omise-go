package operations_test

import (
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestLink_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// create
	link := &omise.Link{}
	client.MustDo(link, &CreateLink{
		Amount:      99900,
		Currency:    "thb",
		Title:       "Hot Latte",
		Description: "A warm cup of coffee",
		Multiple:    true,
	})
	if !a.NotNil(t, link) {
		return
	}

	t.Log("created link:", link.ID)
	a.Equal(t, int64(99900), link.Amount)
	a.Equal(t, "Hot Latte", link.Title)
	a.True(t, link.Multiple)

	// retrieve created link
	link2 := &omise.Link{}
	client.MustDo(link2, &RetrieveLink{LinkID: link.ID})
	a.Equal(t, link.Amount, link2.Amount)
	a.Equal(t, link.Title, link2.Title)

	// list created links from the last hour
	links, list := &omise.LinkList{}, &ListLinks{
		List{Limit: 100, From: time.Now().Add(-1 * time.Hour)},
	}
	client.MustDo(links, list)

	a.True(t, len(links.Data) > 0, "links list empty!")
	link2 = links.Find(link.ID)
	if !a.NotNil(t, link2, "could not find recent links in list.") {
		return
	}

	a.Equal(t, link.ID, link2.ID)
	a.Equal(t, link.Amount, link2.Amount)
	a.Equal(t, link.Title, link2.Title)
}
