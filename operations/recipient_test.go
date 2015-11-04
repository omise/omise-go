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

func TestRecipient(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	// create a recipient
	// sample from: https://www.omise.co/bank-account-api
	bankAccount := &omise.BankAccount{
		Brand:  "bbl",
		Number: "1234567890",
		Name:   "Somchai Prasert",
	}

	create := &CreateRecipient{
		Name:        "Jun Hasegawa",
		Email:       "jun@omise.co",
		Description: "Owns Omise",
		Type:        omise.Individual,
		BankAccount: bankAccount,
	}

	jun := &omise.Recipient{}
	if e := client.Do(jun, create); !(a.NoError(t, e) && a.NotNil(t, jun)) {
		return
	}

	t.Log("created recipient:", jun.ID)
	a.Equal(t, create.Email, jun.Email)
	if a.NotNil(t, jun.Description) {
		a.Equal(t, create.Description, *jun.Description)
	}
	a.Equal(t, jun.BankAccount.Name, bankAccount.Name)

	// list created customers
	recipients, list := &omise.RecipientList{}, &ListRecipients{
		List{From: time.Now().Add(-1 * time.Hour), Limit: 100},
	}
	if e := client.Do(recipients, list); !a.NoError(t, e) {
		return
	}

	a.True(t, len(recipients.Data) > 0, "no created customers in list!")

	jim := recipients.Find(jun.ID)
	a.Equal(t, jun.ID, jim.ID)
	a.Equal(t, jun.Email, jim.Email)

	// // update
	update := &UpdateRecipient{
		RecipientID: jim.ID,
		Description: "I'm JONES now.",
	}

	jones := &omise.Recipient{}
	if e := client.Do(jones, update); !a.NoError(t, e) {
		return
	}

	a.Equal(t, jim.ID, jones.ID)
	if a.NotNil(t, jones.Description) {
		a.Equal(t, update.Description, *jones.Description)
	}

	// fetch
	josh, retreive := &omise.Recipient{}, &RetreiveRecipient{jones.ID}
	if e := client.Do(josh, retreive); !a.NoError(t, e) {
		return
	}

	a.Equal(t, jones.ID, josh.ID)
	a.Equal(t, jones.Email, josh.Email)
	a.Equal(t, jones.Description, josh.Description)

	// delete
	del, destroy := &omise.Deletion{}, &DestroyRecipient{jones.ID}
	if e := client.Do(del, destroy); !a.NoError(t, e) {
		return
	}

	a.Equal(t, jones.Object, del.Object)
	a.Equal(t, jones.ID, del.ID)
	a.Equal(t, jones.Live, del.Live)
	a.True(t, del.Deleted)
}
