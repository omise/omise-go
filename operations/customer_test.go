package operations_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

var _ = fmt.Println

func TestCustomer(t *testing.T) {
	const (
		CustomerID = "cust_test_4yq6txdpfadhbaqnwp3"
		CardID     = "card_test_4yq6tuucl9h4erukfl0"
	)

	client := testutil.NewFixedClient(t)

	customer := &omise.Customer{}
	client.MustDo(customer, &CreateCustomer{})
	a.Equal(t, CustomerID, customer.ID)

	customer = &omise.Customer{}
	client.MustDo(customer, &RetrieveCustomer{CustomerID})
	a.Equal(t, CustomerID, customer.ID)
	a.Equal(t, CardID, customer.DefaultCard)
	if a.Len(t, customer.Cards.Data, 1) {
		a.Equal(t, CardID, customer.Cards.Data[0].ID)
	}

	customers := &omise.CustomerList{}
	client.MustDo(customers, &ListCustomers{})
	if a.Len(t, customers.Data, 1) {
		a.Equal(t, CustomerID, customers.Data[0].ID)
	}

	client.MustDo(customer, &UpdateCustomer{
		CustomerID: customer.ID,
		Email:      "john.doe.the.second@example.com",
	})
	a.Equal(t, "john.doe.the.second@example.com", customer.Email)

	del := &omise.Deletion{}
	client.MustDo(del, &DestroyCustomer{CustomerID})
	a.Equal(t, CustomerID, del.ID)

	e := client.Do(nil, &RetrieveCustomer{"not_exist"})
	if a.Error(t, e) {
		a.EqualError(t, e, "(404/not_found) customer missing was not found")
	}
}

func TestCustomer_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	token := &omise.Token{}
	client.MustDo(token, CreateTokenOp)

	// create a customer
	jack := &omise.Customer{}
	client.MustDo(jack, &CreateCustomer{
		Email:       "chakrit@omise.co",
		Description: "I'm JACK",
		Card:        token.ID,
	})
	if !a.NotNil(t, jack) {
		return
	}

	t.Log("created customer:", jack.ID)
	a.Equal(t, "chakrit@omise.co", jack.Email)
	a.Equal(t, "I'm JACK", jack.Description)
	a.Len(t, jack.Cards.Data, 1)

	// list created customers
	customers := &omise.CustomerList{}
	client.MustDo(customers, &ListCustomers{
		List{
			From:  time.Now().Add(-1 * time.Hour),
			Limit: 100,
		},
	})

	a.True(t, len(customers.Data) > 0, "no created customers in list!")

	jane := customers.Find(jack.ID)
	a.Equal(t, jack.ID, jane.ID)
	a.Equal(t, jack.Email, jane.Email)

	// update
	john := &omise.Customer{}
	client.MustDo(john, &UpdateCustomer{
		CustomerID:  jack.ID,
		Description: "I'm JOHN now.",
	})

	a.Equal(t, jack.ID, john.ID)
	a.Equal(t, "I'm JOHN now.", john.Description)

	// fetch
	jill, retrieve := &omise.Customer{}, &RetrieveCustomer{john.ID}
	client.MustDo(jill, retrieve)

	a.Equal(t, john.ID, jill.ID)
	a.Equal(t, john.Email, jill.Email)
	a.Equal(t, john.Description, jill.Description)

	// delete
	del, destroy := &omise.Deletion{}, &DestroyCustomer{CustomerID: jill.ID}
	client.MustDo(del, destroy)

	a.Equal(t, jill.Object, del.Object)
	a.Equal(t, jill.ID, del.ID)
	a.Equal(t, jill.Live, del.Live)
	a.True(t, del.Deleted)
}
