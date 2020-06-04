package operations


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
	Base
	CustomerID string `json:"-"`
}

func (req *DestroyCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      DELETE,
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
	Base
	CustomerID string `json:"-"`
}

func (req *RetrieveCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
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
	Base
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
		Method:      PATCH,
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
	Base
}

func (req *SearchCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
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
	Base
}

func (req *ListCustomers) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
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
	Base
	Card string `json:"card"`
	Description string `json:"description"`
	Email string `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (req *CreateCustomer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/customers",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Schedule{}, &ListCustomersSchedules{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListCustomersSchedules struct {
	Base
	CustomerID string `json:"-"`
}

func (req *ListCustomersSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/customers/" + req.CustomerID + "/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Card{}, &DestroyCustomerCard{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyCustomerCard struct {
	Base
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

func (req *DestroyCustomerCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      DELETE,
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Card{}, &RetrieveCustomerCard{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveCustomerCard struct {
	Base
	CustomerID string `json:"-"`
	CardID string `json:"-"`
}

func (req *RetrieveCustomerCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Card{}, &UpdateCustomerCard{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateCustomerCard struct {
	Base
	CustomerID string `json:"-"`
	CardID string `json:"-"`
	City string `json:"city"`
	ExpirationMonth int `json:"expiration_month"`
	ExpirationYear int `json:"expiration_year"`
	Name string `json:"name"`
	PostalCode string `json:"postal_code"`
}

func (req *UpdateCustomerCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      PATCH,
		Path:        "/customers/" + req.CustomerID + "/cards/" + req.CardID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Card{}, &ListCustomersCard{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListCustomersCard struct {
	Base
	CustomerID string `json:"-"`
}

func (req *ListCustomersCard) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/customers/" + req.CustomerID + "/cards",
		ContentType: "application/json",
	}
}

