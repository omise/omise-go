package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	transactions, list := &omise.TransactionList{}, &ListTransaction{
//		List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(transactions, list); !a.NoError(t, e) {
//		return
//	}
//
//	fmt.Println("# of transactions in the last hour:", len(transactions.Data))
//
type ListTransactions struct {
	List
}

func (req *ListTransactions) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transactions",
		ContentType: "application/json",
	}
}

// Example:
//
//	transaction, retrieve := &omise.Transaction{}, &RetrieveTransaction{"trxn_987"}
//		TransactionID: transactions.Data[0].ID,
//	}
//	if e := client.Do(transaction, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transaction #trxn_987: %#v\n", transaction)
//
type RetrieveTransaction struct {
	TransactionID string `json:"-"`
}

func (req *RetrieveTransaction) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transactions/" + req.TransactionID,
		ContentType: "application/json",
	}
}
