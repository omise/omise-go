package omise

import (
	"time"
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Transaction represents Transaction object.
// See https://www.omise.co/transactions-api for more information.
type Transaction struct {
	Base
	Amount int64 `json:"amount"`
	Currency string `json:"currency"`
	Direction TransactionDirection `json:"direction"`
	Key string `json:"key"`
	Location string `json:"location"`
	Origin string `json:"origin"`
	TransferableAt time.Time `json:"transferable_at"`
}

// TransactionService represents resource service.
type TransactionService struct {
	*Client
}

// Transaction defines resource service.
func (c *Client) Transaction() *TransactionService {
	return &TransactionService{c}
}

// Retrieve retrieves transaction
//
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
func (s *TransactionService) Retrieve(ctx context.Context, params *RetrieveTransactionParams) (*Transaction, error) {
	result := &Transaction{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveTransactionParams params object.
type RetrieveTransactionParams struct {
	TransactionID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveTransactionParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transactions/" + req.TransactionID,
		ContentType: "application/json",
	}
}

// List lists transactions
//
// Example:
//
//	transaction, list := &omise.TransactionList{}, &ListTransactions{
//	}
//	if e := client.Do(transaction, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("transaction: %#v\n", transaction)
//
func (s *TransactionService) List(ctx context.Context, params *ListTransactionsParams) (*TransactionList, error) {
	result := &TransactionList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListTransactionsParams params object.
type ListTransactionsParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListTransactionsParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/transactions",
		ContentType: "application/json",
	}
}

