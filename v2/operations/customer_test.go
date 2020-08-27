package operations_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

var _ = fmt.Println

func TestCustomer(t *testing.T) {
	const (
		CustomerID = "cust_test_4yq6txdpfadhbaqnwp3"
		CardID     = "card_test_4yq6tuucl9h4erukfl0"
	)

	client := testutil.NewFixedClient(t)

	customer, _ := client.Customer().Create(context.Background(), &omise.CreateCustomerParams{})
	r.Equal(t, CustomerID, customer.ID)

	customer, _ = client.Customer().Retrieve(context.Background(), &omise.RetrieveCustomerParams{CustomerID})
	r.Equal(t, CustomerID, customer.ID)
	r.Equal(t, CardID, customer.DefaultCard)
	r.Len(t, customer.Cards.Data, 1)
	r.Equal(t, CardID, customer.Cards.Data[0].ID)

	customers, _ := client.Customer().List(context.Background(), &omise.ListCustomersParams{})
	r.Len(t, customers.Data, 1)
	r.Equal(t, CustomerID, customers.Data[0].ID)

	customer, _ = client.Customer().Update(context.Background(), &omise.UpdateCustomerParams{
		CustomerID: customer.ID,
		Email:      "john.doe.the.second@example.com",
	})
	r.Equal(t, "john.doe.the.second@example.com", customer.Email)

	del, _ := client.Customer().Destroy(context.Background(), &omise.DestroyCustomerParams{CustomerID})
	r.Equal(t, CustomerID, del.ID)

	_, err := client.Customer().Retrieve(context.Background(), &omise.RetrieveCustomerParams{"not_exist"})
	r.Error(t, err)
	r.EqualError(t, err, "(404/not_found) customer missing was not found")
}

func TestCustomer_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	token := createTestToken(client)

	// create a customer
	jack, _ := client.Customer().Create(context.Background(), &omise.CreateCustomerParams{
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
	customers, _ := client.Customer().List(context.Background(), &omise.ListCustomersParams{
		omise.ListParams{
			From:  time.Now().Add(-1 * time.Hour),
			Limit: 100,
		},
	})

	r.True(t, len(customers.Data) > 0, "no created customers in list!")

	jane := customers.Find(jack.ID)
	r.Equal(t, jack.ID, jane.ID)
	r.Equal(t, jack.Email, jane.Email)

	// update
	john, _ := client.Customer().Update(context.Background(), &omise.UpdateCustomerParams{
		CustomerID:  jack.ID,
		Description: "I'm JOHN now.",
	})

	r.Equal(t, jack.ID, john.ID)
	r.Equal(t, "I'm JOHN now.", john.Description)

	// fetch
	jill, _ := client.Customer().Retrieve(context.Background(), &omise.RetrieveCustomerParams{john.ID})

	r.Equal(t, john.ID, jill.ID)
	r.Equal(t, john.Email, jill.Email)
	r.Equal(t, john.Description, jill.Description)

	// delete
	del, _ := client.Customer().Destroy(context.Background(), &omise.DestroyCustomerParams{CustomerID: jill.ID})

	r.Equal(t, jill.Object, del.Object)
	r.Equal(t, jill.ID, del.ID)
	r.Equal(t, jill.Live, del.Live)
}

func TestListCustomerChargeSchedules(t *testing.T) {
	client := testutil.NewFixedClient(t)
	schds, _ := client.Customer().ListSchedules(context.Background(), &omise.ListCustomerSchedulesParams{
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
	list := &omise.ListCustomerSchedulesParams{
		CustomerID: "cust_1234",
		ListParams: omise.ListParams{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}

	schds, _ := client.Customer().ListSchedules(context.Background(), list)

	t.Logf("Schedules Len: %d\n", len(schds.Data))
	t.Logf("%#v\n", schds)
}

func TestCreateCustomerMarshal_WithMetadata(t *testing.T) {
	req := &CreateCustomer{
		Email:       "john.doe@example.com",
		Description: "John Doe (id: 30)",
		Card:        "card_test_4yq6tuucl9h4erukfl0",
		Metadata: map[string]interface{}{
			"answer": 42,
		},
	}

	expected := `{"card":"card_test_4yq6tuucl9h4erukfl0","description":"John Doe (id: 30)","email":"john.doe@example.com","metadata":{"answer":42}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "error should be nothing")
	r.Equal(t, expected, string(b))
}

func TestCreateCustomerMarshal_WithoutMetadata(t *testing.T) {
	req := &CreateCustomer{
		Email:       "john.doe@example.com",
		Description: "John Doe (id: 30)",
		Card:        "card_test_4yq6tuucl9h4erukfl0",
	}

	expected := `{"card":"card_test_4yq6tuucl9h4erukfl0","description":"John Doe (id: 30)","email":"john.doe@example.com"}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}

func TestUpdateCustomerMarshal_WithMetadata(t *testing.T) {
	req := &UpdateCustomer{
		CustomerID:  "cust_test_4yq6txdpfadhbaqnwp3",
		Email:       "john.doe@example.com",
		Description: "John Doe (id: 30)",
		Card:        "card_test_4yq6tuucl9h4erukfl0",
		Metadata: map[string]interface{}{
			"answer": 42,
		},
	}

	expected := `{"card":"card_test_4yq6tuucl9h4erukfl0","description":"John Doe (id: 30)","email":"john.doe@example.com","metadata":{"answer":42}}`

	b, err := json.Marshal(req)
	r.Nil(t, err, "err should be nothing")
	r.Equal(t, expected, string(b))
}
