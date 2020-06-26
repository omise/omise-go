package omise

import (
	"github.com/omise/omise-go/v2/internal"
	"context"
)

// Customer represents Customer object.
// See https://www.omise.co/customers-api for more information.
type Customer struct {
	Base
	Cards *CardList `json:"cards"`
	DefaultCard string `json:"default_card"`
	Description string `json:"description"`
	Email string `json:"email"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
}

// CustomerService represents resource service.
type CustomerService struct {
	*Client
}

// Customer defines resource service.
func (c *Client) Customer() *CustomerService {
	return &CustomerService{c}
}

// Destroy destroys customer
//
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
func (s *CustomerService) Destroy(ctx context.Context, params *DestroyCustomerParams) (*Customer, error) {
	result := &Customer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DestroyCustomerParams params object.
type DestroyCustomerParams struct {
	CustomerID string `json:"-"`
}

// Describe describes structure of request
func (req *DestroyCustomerParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "DELETE",
		Path:        "/customers/" + req.CustomerID,
		ContentType: "application/json",
	}
}

// Retrieve retrieves customer
//
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
func (s *CustomerService) Retrieve(ctx context.Context, params *RetrieveCustomerParams) (*Customer, error) {
	result := &Customer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RetrieveCustomerParams params object.
type RetrieveCustomerParams struct {
	CustomerID string `json:"-"`
}

// Describe describes structure of request
func (req *RetrieveCustomerParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID,
		ContentType: "application/json",
	}
}

// Update updates customer
//
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
func (s *CustomerService) Update(ctx context.Context, params *UpdateCustomerParams) (*Customer, error) {
	result := &Customer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// UpdateCustomerParams params object.
type UpdateCustomerParams struct {
	CustomerID string `json:"-"`
	Card string `json:"card,omitempty"`
	DefaultCard string `json:"default_card,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe describes structure of request
func (req *UpdateCustomerParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "PATCH",
		Path:        "/customers/" + req.CustomerID,
		ContentType: "application/json",
	}
}

// List lists customers
//
// Example:
//
//	customer, list := &omise.CustomerList{}, &ListCustomers{
//	}
//	if e := client.Do(customer, list); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("customer: %#v\n", customer)
//
func (s *CustomerService) List(ctx context.Context, params *ListCustomersParams) (*CustomerList, error) {
	result := &CustomerList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListCustomersParams params object.
type ListCustomersParams struct {
	ListParams
}

// Describe describes structure of request
func (req *ListCustomersParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers",
		ContentType: "application/json",
	}
}

// Create creates customer
//
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
func (s *CustomerService) Create(ctx context.Context, params *CreateCustomerParams) (*Customer, error) {
	result := &Customer{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CreateCustomerParams params object.
type CreateCustomerParams struct {
	Card string `json:"card,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Describe describes structure of request
func (req *CreateCustomerParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "POST",
		Path:        "/customers",
		ContentType: "application/json",
	}
}

// ListSchedules lists customer schedules
//
// Example:
//
//	schedule, listSchedules := &omise.ScheduleList{}, &ListCustomerSchedules{
//		CustomerID: "cust_8b6jt4hr5i"
//	}
//	if e := client.Do(schedule, listSchedules); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("schedule: %#v\n", schedule)
//
func (s *CustomerService) ListSchedules(ctx context.Context, params *ListCustomerSchedulesParams) (*ScheduleList, error) {
	result := &ScheduleList{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ListCustomerSchedulesParams params object.
type ListCustomerSchedulesParams struct {
	CustomerID string `json:"-"`
	ListParams
}

// Describe describes structure of request
func (req *ListCustomerSchedulesParams) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/schedules",
		ContentType: "application/json",
	}
}

