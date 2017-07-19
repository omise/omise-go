package operations

import (
	"encoding/json"

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
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/customers",
	}
}

// Example:
//
//	customer, create := &omise.Customer{}, &CreateCustomer{
//		Email:       "kokyaku@omise.co",
//		Description: "I'm a customer",
//		Card:        token.ID,
//	})
//	if e := client.Do(customer, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("created customer: %#v\n", customer)
//
type CreateCustomer struct {
	Email       string
	Description string
	Card        string
}

func (req *CreateCustomer) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "POST",
		Path:     "/customers",
	}
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
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "GET",
		Path:     "/customers/" + req.CustomerID,
	}
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
	DefaultCard string `json:"default_card"`
}

func (req *UpdateCustomer) Op() *internal.Op {
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "PATCH",
		Path:     "/customers/" + req.CustomerID,
	}
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
	return &internal.Op{
		Endpoint: internal.API,
		Method:   "DELETE",
		Path:     "/customers/" + req.CustomerID,
	}
}

// ListCustomerChargeSchedules represent list customer schedule schedules API payload
//
// Example:
//
//      var schds omise.ScheduleList
//	list := ListCustomerChargeSchedules{
//		CustomerID: "cust_123"
//		List: List{
//			Limit: 100,
//			From: time.Now().Add(-1 * time.Hour),
//		},
//	}
//	if e := client.Do(&schds, &list); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("# of schedules made in the last hour:", len(schds.Data))
//
type ListCustomerChargeSchedules struct {
	CustomerID string
	List
}

func (req *ListCustomerChargeSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(req.List)
}

func (req *ListCustomerChargeSchedules) Op() *internal.Op {
	return &internal.Op{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/schedules",
		ContentType: "application/json",
	}
}
