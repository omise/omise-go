package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	customer, destroy := &omise.Customer{}, &DestroyCustomer{
//		CustomerID: "cust_8b6jt4hr5i"
//	}
//	if e := client.Do(customer, destroy); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("customer: %#v\n", customer)
//
type DestroyCustomer struct {
	CustomerID string `json:"-"`
}

// Describe
func (req *DestroyCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/customers/" + req.CustomerID,
		ContentType: "application/json",
	}
}

// Example:
//
//	customer, retrieve := &omise.Customer{}, &RetrieveCustomer{
//		CustomerID: "cust_8b6jt4hr5i"
//	}
//	if e := client.Do(customer, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("customer: %#v\n", customer)
//
type RetrieveCustomer struct {
	CustomerID string `json:"-"`
}

// Describe
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
//	customer, update := &omise.Customer{}, &UpdateCustomer{
//		CustomerID: "cust_8b6jt4hr5i"
//	}
//	if e := client.Do(customer, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("customer: %#v\n", customer)
//
type UpdateCustomer struct {
	CustomerID string `json:"-"`
	Card string `json:"card,omitempty"`
	DefaultCard string `json:"default_card,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe
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
//	customer, list := &omise.Customer{}, &ListCustomers{
//	}
//	if e := client.Do(customer, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("customer: %#v\n", customer)
//
type ListCustomers struct {
	List
}

// Describe
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
//	}
//	if e := client.Do(customer, create); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("customer: %#v\n", customer)
//
type CreateCustomer struct {
	Card string `json:"card,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe
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
//	schedule, listSchedules := &omise.Schedule{}, &ListCustomerSchedules{
//		CustomerID: "cust_8b6jt4hr5i"
//	}
//	if e := client.Do(schedule, listSchedules); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
type ListCustomerSchedules struct {
	CustomerID string `json:"-"`
	List
}

// Describe
func (req *ListCustomerSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/schedules",
		ContentType: "application/json",
	}
}

