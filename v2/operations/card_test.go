package operations_test

import (
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

const (
	CustomerID = "cust_test_4yq6txdpfadhbaqnwp3"
	CardID     = "card_test_4yq6tuucl9h4erukfl0"
)

func TestCard(t *testing.T) {
	client := testutil.NewFixedClient(t)
	card := &omise.Card{}
	client.MustDo(card, &RetrieveCard{CustomerID, CardID})
	r.Equal(t, CardID, card.ID)

	card = &omise.Card{}
	err := client.Do(card, &RetrieveCard{CustomerID, "not_exist"})
	r.Error(t, err)
	r.EqualError(t, err, "(404/not_found) customer missing was not found")

	client.MustDo(card, &UpdateCard{
		CustomerID: CustomerID,
		CardID:     CardID,
		Name:       "JOHN W. DOE",
	})
	r.Equal(t, "JOHN W. DOE", card.Name)

	del := &omise.Deletion{}
	client.MustDo(del, &DestroyCard{CustomerID, CardID})
	r.Equal(t, card.ID, del.ID)
}

func TestCard_Network(t *testing.T) {
	testutil.Require(t, "network")
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
	r.NotNil(t, cards)
	r.Len(t, cards.Data, 0)

	// add some cards
	tok1, tok2 := createTestToken(client), createTestToken(client)
	client.MustDo(customer, &UpdateCustomer{CustomerID: customer.ID, Card: tok1.ID})
	client.MustDo(customer, &UpdateCustomer{CustomerID: customer.ID, Card: tok2.ID})

	// see added cards (using pagination API to call twice)
	firstHalf, secondHalf := &omise.CardList{}, &omise.CardList{}
	client.MustDo(firstHalf, &ListCards{customer.ID, List{Offset: 0, Limit: 1}})
	client.MustDo(secondHalf, &ListCards{customer.ID, List{Offset: 1, Limit: 1}})
	r.Len(t, firstHalf.Data, 1)
	r.Len(t, secondHalf.Data, 1)
	r.NotEqual(t, firstHalf.Data[0].ID, secondHalf.Data[0].ID)

	cards = &omise.CardList{
		List: omise.List{Offset: 0, Limit: 0},
		Data: append(firstHalf.Data, secondHalf.Data...),
	}
	r.Len(t, cards.Data, 2)

	// update a card
	card, card2 := cards.Data[0], &omise.Card{}
	client.MustDo(card2, &UpdateCard{
		CustomerID: customer.ID,
		CardID:     card.ID,
		Name:       "Changed my name",
	})

	r.Equal(t, card.ID, card2.ID)
	r.Equal(t, "Changed my name", card2.Name)

	// retrieve added cards
	card = &omise.Card{}
	client.MustDo(card, &RetrieveCard{CustomerID: customer.ID, CardID: cards.Data[0].ID})

	r.Equal(t, cards.Data[0].ID, card.ID)
	r.Equal(t, cards.Data[0].LastDigits, card.LastDigits)

	// remove cards
	del1, del2 := &omise.Deletion{}, &omise.Deletion{}
	client.MustDo(del1, &DestroyCard{CustomerID: customer.ID, CardID: cards.Data[0].ID})
	client.MustDo(del2, &DestroyCard{CustomerID: customer.ID, CardID: cards.Data[1].ID})

	r.Equal(t, cards.Data[0].ID, del1.ID)
	r.Equal(t, cards.Data[1].ID, del2.ID)
	r.True(t, del1.Deleted)
	r.True(t, del2.Deleted)

	// list should be empty again
	client.MustDo(cards, &ListCards{CustomerID: customer.ID})
	r.Len(t, cards.Data, 0)
}
