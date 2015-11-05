package operations

import (
	"github.com/omise/omise-go/internal"
)

// Example:
//
//	customers, list := &omise.CustomerList{}, &ListCustomers{
//		List{
//			From: time.Now().Add(-1 * time.Hour),
//			Limit: 100
//		},
//	}
//	if e := client.Do(customers, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of new customers in the last hour:", len(customers.Data))
//
type ListCustomers struct {
	List
}

func (req *ListCustomers) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers", nil}
}

type CreateCustomer struct {
	Email       string
	Description string
	Card        string
}

func (req *CreateCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "POST", "/customers", nil}
}

// Example:
//
//	cust, retrieve := &omise.Customer{}, &RetrieveCustomer{"cust_123"}
//	if e := client.Do(cust, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("cust_123: %#v\n", cust)
//
type RetrieveCustomer struct {
	CustomerID string `query:"-"`
}

func (req *RetrieveCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "GET", "/customers/" + req.CustomerID, nil}
}

// Example:
//
//	cust, update := &&omise.Customer{}, UpdateCustomer{
//		CustomerID:  "cust_987",
//		Description: "I'm JOHN now.",
//	}
//	if e := client.Do(cust, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated customer: %#v\n", cust)
//
type UpdateCustomer struct {
	CustomerID  string `query:"-"`
	Email       string
	Description string
	Card        string
}

func (req *UpdateCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "PATCH", "/customers/" + req.CustomerID, nil}
}

// Example:
//
//	del, destroy := &omise.Deletion{}, &DestroyCustomer{
//		CustomerID: "cust_123",
//	}
//	if e := client.Do(del, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("destroyed customer:", del.ID)
//
type DestroyCustomer struct {
	CustomerID string `query:"-"`
}

func (req *DestroyCustomer) Op() *internal.Op {
	return &internal.Op{internal.API, "DELETE", "/customers/" + req.CustomerID, nil}
}
