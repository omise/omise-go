package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	transaction, retrieve := &omise.Transaction{}, &RetrieveTransaction{
//		TransactionID: "trxn_5is3gq69cm"
//	}
//	if e := client.Do(transaction, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transaction: %#v\n", transaction)
//
type RetrieveTransaction struct {
	TransactionID string `json:"-"`
}

// Describe
func (req *RetrieveTransaction) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transactions/" + req.TransactionID,
		ContentType: "application/json",
	}
}

// Example:
//
//	transaction, list := &omise.Transaction{}, &ListTransactions{
//	}
//	if e := client.Do(transaction, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transaction: %#v\n", transaction)
//
type ListTransactions struct {
	List
}

// Describe
func (req *ListTransactions) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transactions",
		ContentType: "application/json",
	}
}

