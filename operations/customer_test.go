package operations_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	r "github.com/stretchr/testify/require"
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
	r.Equal(t, CustomerID, customer.ID)

	customer = &omise.Customer{}
	client.MustDo(customer, &RetrieveCustomer{CustomerID})
	r.Equal(t, CustomerID, customer.ID)
	r.Equal(t, CardID, customer.DefaultCard)
	r.Len(t, customer.Cards.Data, 1)
	r.Equal(t, CardID, customer.Cards.Data[0].ID)

	customers := &omise.CustomerList{}
	client.MustDo(customers, &ListCustomers{})
	r.Len(t, customers.Data, 1)
	r.Equal(t, CustomerID, customers.Data[0].ID)

	client.MustDo(customer, &UpdateCustomer{
		CustomerID: customer.ID,
		Email:      "john.doe.the.second@example.com",
	})
	r.Equal(t, "john.doe.the.second@example.com", customer.Email)

	del := &omise.Deletion{}
	client.MustDo(del, &DestroyCustomer{CustomerID})
	r.Equal(t, CustomerID, del.ID)

	e := client.Do(nil, &RetrieveCustomer{"not_exist"})
	r.Error(t, e)
	r.EqualError(t, e, "(404/not_found) customer missing was not found")
}

func TestCustomer_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	// create a customer
	jack := &omise.Customer{}
	client.MustDo(jack, &CreateCustomer{
		Email:       "chakrit@omise.co",
		Description: "I'm JACK",
		Card:        token.ID,
	})
	r.NotNil(t, jack)

	t.Log("created customer:", jack.ID)
	r.Equal(t, "chakrit@omise.co", jack.Email)
	r.Equal(t, "I'm JACK", jack.Description)
	r.Len(t, jack.Cards.Data, 1)

	// list created customers
	customers := &omise.CustomerList{}
	client.MustDo(customers, &ListCustomers{
		List{
			From:  time.Now().Add(-1 * time.Hour),
			Limit: 100,
		},
	})

	r.True(t, len(customers.Data) > 0, "no created customers in list!")

	jane := customers.Find(jack.ID)
	r.Equal(t, jack.ID, jane.ID)
	r.Equal(t, jack.Email, jane.Email)

	// update
	john := &omise.Customer{}
	client.MustDo(john, &UpdateCustomer{
		CustomerID:  jack.ID,
		Description: "I'm JOHN now.",
	})

	r.Equal(t, jack.ID, john.ID)
	r.Equal(t, "I'm JOHN now.", john.Description)

	// fetch
	jill, retrieve := &omise.Customer{}, &RetrieveCustomer{john.ID}
	client.MustDo(jill, retrieve)

	r.Equal(t, john.ID, jill.ID)
	r.Equal(t, john.Email, jill.Email)
	r.Equal(t, john.Description, jill.Description)

	// delete
	del, destroy := &omise.Deletion{}, &DestroyCustomer{CustomerID: jill.ID}
	client.MustDo(del, destroy)

	r.Equal(t, jill.Object, del.Object)
	r.Equal(t, jill.ID, del.ID)
	r.Equal(t, jill.Live, del.Live)
	r.True(t, del.Deleted)
}

func TestListCustomerChargeSchedules(t *testing.T) {
	client := testutil.NewFixedClient(t)
	var schds omise.ScheduleList
	client.MustDo(&schds, &ListCustomerChargeSchedules{
		CustomerID: "cust_test_4yq6txdpfadhbaqnwp3",
	})

	r.Len(t, schds.Data, 1)

	r.Equal(t, "schd_57zhl296uxc7yiun6xa", schds.Data[0].ID)
	r.NotNil(t, schds.Data[0].Charge)
	r.Nil(t, schds.Data[0].Transfer)
}

func TestListCustomerChargeSchedules_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	var schds omise.ScheduleList
	list := ListCustomerChargeSchedules{
		CustomerID: "cust_1234",
		List: List{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}

	client.MustDo(&schds, &list)

	t.Logf("Schedules Len: %d\n", len(schds.Data))
	t.Logf("%#v\n", schds)
}
