package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	create := &CreateCustomer{
		Email:       "chakrit@omise.co",
		Description: "Omise",
		Card:        "",
	}

	customer := &omise.Customer{}
	if e := client.Do(customer, create); !a.NoError(t, e) {
		return
	}

	// list empty collection
	list := &ListCards{CustomerID: customer.ID}
	cards := &omise.CardList{}
	if e := client.Do(cards, list); !a.NoError(t, e) {
		return
	}

	if !a.NotNil(t, cards) {
		return
	}

	a.Len(t, cards.Data, 0)

	// add some cards
	tok1, tok2 := &omise.Token{}, &omise.Token{}
	if e := client.Do(tok1, CreateTokenOp); !a.NoError(t, e) {
		return
	} else if e := client.Do(tok2, CreateTokenOp); !a.NoError(t, e) {
		return
	}

	update1 := &UpdateCustomer{CustomerID: customer.ID, Card: tok1.ID}
	update2 := &UpdateCustomer{CustomerID: customer.ID, Card: tok2.ID}
	if e := client.Do(customer, update1); !a.NoError(t, e) {
		return
	} else if e := client.Do(customer, update2); !a.NoError(t, e) {
		return
	}

	// see added cards
	if e := client.Do(cards, list); !a.NoError(t, e) {
		return
	}

	a.Len(t, cards.Data, 2)

	// retreive added cards
	card := &omise.Card{}
	retreive := &RetreiveCard{CustomerID: customer.ID, CardID: cards.Data[0].ID}
	if e := client.Do(card, retreive); !a.NoError(t, e) {
		return
	}

	a.Equal(t, cards.Data[0].ID, card.ID)
	a.Equal(t, cards.Data[0].LastDigits, card.LastDigits)

	// remove cards
	del1, del2 := &omise.Deletion{}, &omise.Deletion{}
	if e := client.Do(del1, &DestroyCard{CustomerID: customer.ID, CardID: cards.Data[0].ID}); !a.NoError(t, e) {
		return
	} else if e := client.Do(del2, &DestroyCard{CustomerID: customer.ID, CardID: cards.Data[1].ID}); !a.NoError(t, e) {
		return
	}

	a.Equal(t, cards.Data[0].ID, del1.ID)
	a.Equal(t, cards.Data[1].ID, del2.ID)
	a.True(t, del1.Deleted)
	a.True(t, del2.Deleted)

	// list should be empty again
	if e := client.Do(cards, list); !a.NoError(t, e) {
		return
	}

	a.Len(t, cards.Data, 0)
}
