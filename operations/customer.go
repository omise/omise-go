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

func (req *ListCustomers) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers",
		ContentType: "application/json",
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
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
	Card        string `json:"card,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func (req *CreateCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/customers",
		ContentType: "application/json",
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
	CustomerID string `json:"-"`
}

func (req *RetrieveCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID,
		ContentType: "application/json",
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
	CustomerID  string `json:"-"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
	Card        string `json:"card,omitempty"`
	DefaultCard string `json:"default_card,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func (req *UpdateCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/customers/" + req.CustomerID,
		ContentType: "application/json",
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
	CustomerID string `json:"-"`
}

func (req *DestroyCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/customers/" + req.CustomerID,
		ContentType: "application/json",
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

func (req *ListCustomerChargeSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/schedules",
		ContentType: "application/json",
	}
}
