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

func TestRecipient(t *testing.T) {
	const (
		RecipientID   = "recp_test_50894vc13y8z4v51iuc"
		BankAccountID = ""
	)

	client := testutil.NewFixedClient(t)

	recipient := &omise.Recipient{}
	client.MustDo(recipient, &RetrieveRecipient{RecipientID})
	r.Equal(t, RecipientID, recipient.ID)
	r.NotNil(t, recipient.BankAccount)
	r.Equal(t, "6789", recipient.BankAccount.LastDigits)

	recipients := &omise.RecipientList{}
	client.MustDo(recipients, &ListRecipients{})
	r.Len(t, recipients.Data, 1)
	r.Equal(t, RecipientID, recipients.Data[0].ID)

	client.MustDo(recipient, &UpdateRecipient{
		RecipientID: RecipientID,
		Email:       "john@doe.com",
	})
	r.Equal(t, "john@doe.com", recipient.Email)

	del := &omise.Deletion{}
	client.MustDo(del, &DestroyRecipient{RecipientID})
	r.Equal(t, recipient.ID, del.ID)
}

func TestRecipient_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	// create a recipient
	// sample from: https://www.omise.co/bank-account-api
	jun, bankAccount := &omise.Recipient{}, &omise.BankAccount{
		Brand:  "bbl",
		Number: "1234567890",
		Name:   "Somchai Prasert",
	}
	client.MustDo(jun, &CreateRecipient{
		Name:        "Jun Hasegawa",
		Email:       "jun@omise.co",
		Description: "Owns Omise",
		Type:        omise.Individual,
		BankAccount: bankAccount,
	})

	t.Log("created recipient:", jun.ID)
	r.Equal(t, "jun@omise.co", jun.Email)
	r.NotNil(t, jun.Description)
	r.Equal(t, "Owns Omise", *jun.Description)
	r.Equal(t, jun.BankAccount.Name, bankAccount.Name)

	// list created customers
	recipients := &omise.RecipientList{}
	client.MustDo(recipients, &ListRecipients{
		List{From: time.Now().Add(-1 * time.Hour), Limit: 100},
	})

	r.True(t, len(recipients.Data) > 0, "no created customers in list!")

	jim := recipients.Find(jun.ID)
	r.Equal(t, jun.ID, jim.ID)
	r.Equal(t, jun.Email, jim.Email)

	// update
	jones := &omise.Recipient{}
	client.MustDo(jones, &UpdateRecipient{
		RecipientID: jim.ID,
		Description: "I'm JONES now.",
	})

	r.Equal(t, jim.ID, jones.ID)
	r.NotNil(t, jones.Description)
	r.Equal(t, "I'm JONES now.", *jones.Description)

	// fetch
	josh := &omise.Recipient{}
	client.MustDo(josh, &RetrieveRecipient{jones.ID})

	r.Equal(t, jones.ID, josh.ID)
	r.Equal(t, jones.Email, josh.Email)
	r.Equal(t, jones.Description, josh.Description)

	// delete
	del := &omise.Deletion{}
	client.MustDo(del, &DestroyRecipient{jones.ID})

	r.Equal(t, jones.Object, del.Object)
	r.Equal(t, jones.ID, del.ID)
	r.Equal(t, jones.Live, del.Live)
	r.True(t, del.Deleted)
}
