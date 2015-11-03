package operations_test

import (
	"fmt"
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

var _ = fmt.Println

func TestCustomer(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	token := &omise.Token{}
	if e := client.Do(token, CreateTokenOp); !a.NoError(t, e) {
		return
	}

	// create a customer
	create := &CreateCustomer{
		Email:       "chakrit@omise.co",
		Description: "I'm JACK",
		Card:        token.ID,
	}

	jack := &omise.Customer{}
	if e := client.Do(jack, create); !a.NoError(t, e) {
		return
	}

	if a.NotNil(t, jack) {
		t.Log("created customer:", jack.ID)
		a.Equal(t, create.Email, jack.Email)
		a.Equal(t, create.Description, jack.Description)
		a.Len(t, jack.Cards.Data, 1)
	} else {
		return
	}

	// update
	update := &UpdateCustomer{
		CustomerID:  jack.ID,
		Description: "I'm JOHN now.",
	}

	john := &omise.Customer{}
	if e := client.Do(john, update); !a.NoError(t, e) {
		return
	}

	a.Equal(t, jack.ID, john.ID)
	a.Equal(t, update.Description, john.Description)

	// fetch
	jill, retreive := &omise.Customer{}, &RetreiveCustomer{john.ID}
	if e := client.Do(jill, retreive); !a.NoError(t, e) {
		return
	}

	a.Equal(t, john.ID, jill.ID)
	a.Equal(t, john.Email, jill.Email)
	a.Equal(t, john.Description, jill.Description)

	// delete
	del, destroy := &omise.Deletion{}, &DestroyCustomer{CustomerID: jill.ID}
	if e := client.Do(del, destroy); !a.NoError(t, e) {
		return
	}

	a.Equal(t, jill.Object, del.Object)
	a.Equal(t, jill.ID, del.ID)
	a.Equal(t, jill.Live, del.Live)
	a.True(t, del.Deleted)
}
