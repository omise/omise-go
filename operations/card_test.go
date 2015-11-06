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

	customer := &omise.Customer{}
	client.MustDo(customer, &CreateCustomer{
		Email:       "chakrit@omise.co",
		Description: "Omise",
		Card:        "",
	})

	defer client.MustDo(nil, &DestroyCustomer{customer.ID})

	// list empty collection
	cards := &omise.CardList{}
	client.MustDo(cards, &ListCards{CustomerID: customer.ID})
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

	client.MustDo(customer, &UpdateCustomer{CustomerID: customer.ID, Card: tok1.ID})
	client.MustDo(customer, &UpdateCustomer{CustomerID: customer.ID, Card: tok2.ID})

	// see added cards (using pagination API to call twice)
	firstHalf, secondHalf := &omise.CardList{}, &omise.CardList{}
	client.MustDo(firstHalf, &ListCards{customer.ID, List{Offset: 0, Limit: 1}})
	client.MustDo(secondHalf, &ListCards{customer.ID, List{Offset: 1, Limit: 1}})

	if !(a.Len(t, firstHalf.Data, 1) && a.Len(t, secondHalf.Data, 1)) {
		return
	} else if !a.NotEqual(t, firstHalf.Data[0].ID, secondHalf.Data[0].ID) {
		return
	}

	cards = &omise.CardList{
		List: omise.List{Offset: 0, Limit: 0},
		Data: append(firstHalf.Data, secondHalf.Data...),
	}
	a.Len(t, cards.Data, 2)

	// update a card
	card, card2 := cards.Data[0], &omise.Card{}
	client.MustDo(card2, &UpdateCard{
		CustomerID: customer.ID,
		CardID:     card.ID,
		Name:       "Changed my name",
	})

	a.Equal(t, card.ID, card2.ID)
	a.Equal(t, "Changed my name", card2.Name)

	// retrieve added cards
	card = &omise.Card{}
	client.MustDo(card, &RetrieveCard{CustomerID: customer.ID, CardID: cards.Data[0].ID})

	a.Equal(t, cards.Data[0].ID, card.ID)
	a.Equal(t, cards.Data[0].LastDigits, card.LastDigits)

	// remove cards
	del1, del2 := &omise.Deletion{}, &omise.Deletion{}
	client.MustDo(del1, &DestroyCard{CustomerID: customer.ID, CardID: cards.Data[0].ID})
	client.MustDo(del2, &DestroyCard{CustomerID: customer.ID, CardID: cards.Data[1].ID})

	a.Equal(t, cards.Data[0].ID, del1.ID)
	a.Equal(t, cards.Data[1].ID, del2.ID)
	a.True(t, del1.Deleted)
	a.True(t, del2.Deleted)

	// list should be empty again
	client.MustDo(cards, &ListCards{CustomerID: customer.ID})
	a.Len(t, cards.Data, 0)
}
