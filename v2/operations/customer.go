package operations

import (
	"github.com/omise/omise-go/v2/internal"
)

// Example:
//
//	charge, update := &omise.Customer{}, &DestroyCustomer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
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

// Example:
//
//	charge, update := &omise.Customer{}, &RetrieveCustomer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
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
//	charge, update := &omise.Customer{}, &UpdateCustomer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateCustomer struct {
	CustomerID string `json:"-"`
	Card string `json:"card"`
	DefaultCard string `json:"default_card"`
	Description string `json:"description"`
	Email string `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
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
//	charge, update := &omise.Customer{}, &SearchCustomer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchCustomer struct {
}

func (req *SearchCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Customer{}, &ListCustomers{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListCustomers struct {
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
//	charge, update := &omise.Customer{}, &CreateCustomer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateCustomer struct {
	Card string `json:"card"`
	Description string `json:"description"`
	Email string `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
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
//	charge, update := &omise.Schedule{}, &ListCustomerSchedules{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListCustomerSchedules struct {
	CustomerID string `json:"-"`
}

func (req *ListCustomerSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      "GET",
		Path:        "/customers/" + req.CustomerID + "/schedules",
		ContentType: "application/json",
	}
}

