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
