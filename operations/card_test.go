package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	client := testutil.NewTestClient(t)

	create := &CreateCustomer{
		Email:       "chakrit@omise.co",
		Description: "Omise",
		Card:        "",
	}

	customer := &omise.Customer{}
	client.MustDo(customer, create)
	defer client.MustDo(nil, &DestroyCustomer{customer.ID})

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

	// see added cards (using pagination API to call twice)
	firstHalf := &omise.CardList{}
	list.Offset = 0
	list.Limit = 1
	if e := client.Do(firstHalf, list); !a.NoError(t, e) {
		return
	}

	a.Len(t, firstHalf.Data, 1)

	secondHalf := &omise.CardList{}
	list.Offset = 1
	list.Limit = 1
	if e := client.Do(secondHalf, list); !a.NoError(t, e) {
		return
	}

	a.Len(t, secondHalf.Data, 1)

	cards = &omise.CardList{
		List: omise.List{Offset: 0, Limit: 0},
		Data: append(firstHalf.Data, secondHalf.Data...),
	}
	a.Len(t, cards.Data, 2)

	// update a card
	card := cards.Data[0]
	card2, update := &omise.Card{}, &UpdateCard{
		CustomerID: customer.ID,
		CardID:     card.ID,
		Name:       "Changed my name",
	}
	if e := client.Do(card2, update); !a.NoError(t, e) {
		return
	}

	a.Equal(t, card2.ID, card.ID)
	a.Equal(t, card2.Name, update.Name)

	// retrieve added cards
	card, retrieve := &omise.Card{}, &RetrieveCard{CustomerID: customer.ID, CardID: cards.Data[0].ID}
	if e := client.Do(card, retrieve); !a.NoError(t, e) {
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
