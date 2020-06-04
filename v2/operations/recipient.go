package operations


// Example:
//
//	charge, update := &omise.Recipient{}, &DestroyRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyRecipient struct {
	Base
	RecipientID string `json:"-"`
}

func (req *DestroyRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      DELETE,
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &RetrieveRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveRecipient struct {
	Base
	RecipientID string `json:"-"`
}

func (req *RetrieveRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &UpdateRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateRecipient struct {
	Base
	RecipientID string `json:"-"`
	BankAccount *BankAccount `json:"bank_account"`
	Description string `json:"description"`
	Email string `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
	Name string `json:"name"`
	TaxID string `json:"tax_id"`
	Type *Type `json:"type"`
}

func (req *UpdateRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      PATCH,
		Path:        "/recipients/" + req.RecipientID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &SearchRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchRecipient struct {
	Base
}

func (req *SearchRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/recipients/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &ListRecipients{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListRecipients struct {
	Base
}

func (req *ListRecipients) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/recipients",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &CreateRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateRecipient struct {
	Base
	BankAccount *BankAccount `json:"bank_account"`
	Description string `json:"description"`
	Email string `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
	Name string `json:"name"`
	TaxID string `json:"tax_id"`
	Type *Type `json:"type"`
}

func (req *CreateRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/recipients",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Schedule{}, &ListRecipientsSchedules{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListRecipientsSchedules struct {
	Base
	RecipientID string `json:"-"`
}

func (req *ListRecipientsSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/recipients/" + req.RecipientID + "/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Recipient{}, &VerifyRecipient{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type VerifyRecipient struct {
	Base
	RecipientID string `json:"-"`
}

func (req *VerifyRecipient) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      PATCH,
		Path:        "/recipients/" + req.RecipientID + "/verify",
		ContentType: "application/json",
	}
}

