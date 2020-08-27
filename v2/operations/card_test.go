package operations_test

import (
	"context"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	_ "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

const (
	CustomerID = "cust_test_4yq6txdpfadhbaqnwp3"
	CardID     = "card_test_4yq6tuucl9h4erukfl0"
)

func TestCard(t *testing.T) {
	client := testutil.NewFixedClient(t)
	card, err := client.Card().Retrieve(context.Background(), &omise.RetrieveCardParams{
		CustomerID: CustomerID,
		CardID:     CardID,
	})

	r.Nil(t, err)
	r.Equal(t, CardID, card.ID)

	card, err = client.Card().Retrieve(context.Background(), &omise.RetrieveCardParams{
		CustomerID: CustomerID,
		CardID:     "not_exist",
	})
	r.Error(t, err)
	r.EqualError(t, err, "(404/not_found) customer missing was not found")

	card, err = client.Card().Update(context.Background(), &omise.UpdateCardParams{
		CustomerID: CustomerID,
		CardID:     CardID,
		Name:       "JOHN W. DOE",
	})
	r.Nil(t, err)
	r.Equal(t, "JOHN W. DOE", card.Name)

	del, err := client.Card().Destroy(context.Background(), &omise.DestroyCardParams{CustomerID, CardID})
	r.Nil(t, err)
	r.Equal(t, card.ID, del.ID)
}

func TestCard_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	customer, err := client.Customer().Create(context.Background(), &omise.CreateCustomerParams{
		Email:       "chakrit@omise.co",
		Description: "Omise",
		Card:        "",
	})
	r.Nil(t, err)

	defer client.Customer().Destroy(context.Background(), &omise.DestroyCustomerParams{customer.ID})

	// list empty collection
	cards, err := client.Card().List(context.Background(), &omise.ListCardsParams{CustomerID: customer.ID})
	r.NotNil(t, cards)
	r.Len(t, cards.Data, 0)

	// add some cards
	tok1, tok2 := createTestToken(client), createTestToken(client)
	client.Customer().Update(context.Background(), &omise.UpdateCustomerParams{CustomerID: customer.ID, Card: tok1.ID})
	client.Customer().Update(context.Background(), &omise.UpdateCustomerParams{CustomerID: customer.ID, Card: tok2.ID})

	// see added cards (using pagination API to call twice)
	firstHalf, _ := client.Card().List(context.Background(), &omise.ListCardsParams{customer.ID, omise.ListParams{Offset: 0, Limit: 1}})
	secondHalf, _ := client.Card().List(context.Background(), &omise.ListCardsParams{customer.ID, omise.ListParams{Offset: 1, Limit: 1}})

	r.Len(t, firstHalf.Data, 1)
	r.Len(t, secondHalf.Data, 1)
	r.NotEqual(t, firstHalf.Data[0].ID, secondHalf.Data[0].ID)

	cards = &omise.CardList{
		List: omise.List{Offset: 0, Limit: 0},
		Data: append(firstHalf.Data, secondHalf.Data...),
	}
	r.Len(t, cards.Data, 2)

	// update a card
	card := cards.Data[0]
	card2, _ := client.Card().Update(context.Background(), &omise.UpdateCardParams{
		CustomerID: customer.ID,
		CardID:     card.ID,
		Name:       "Changed my name",
	})

	r.Equal(t, card.ID, card2.ID)
	r.Equal(t, "Changed my name", card2.Name)

	// retrieve added cards
	card, _ = client.Card().Retrieve(context.Background(), &omise.RetrieveCardParams{CustomerID: customer.ID, CardID: cards.Data[0].ID})

	r.Equal(t, cards.Data[0].ID, card.ID)
	r.Equal(t, cards.Data[0].LastDigits, card.LastDigits)

	// remove cards
	del1, _ := client.Card().Destroy(context.Background(), &omise.DestroyCardParams{CustomerID: customer.ID, CardID: cards.Data[0].ID})
	del2, _ := client.Card().Destroy(context.Background(), &omise.DestroyCardParams{CustomerID: customer.ID, CardID: cards.Data[1].ID})

	r.Equal(t, cards.Data[0].ID, del1.ID)
	r.Equal(t, cards.Data[1].ID, del2.ID)

	// list should be empty again
	cards, _ = client.Card().List(context.Background(), &omise.ListCardsParams{CustomerID: customer.ID})
	r.Len(t, cards.Data, 0)
}
