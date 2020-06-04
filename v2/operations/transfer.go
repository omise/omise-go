package operations


// Example:
//
//	charge, update := &omise.Transfer{}, &DestroyTransfer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type DestroyTransfer struct {
	Base
	TransferID string `json:"-"`
}

func (req *DestroyTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      DELETE,
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transfer{}, &RetrieveTransfer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type RetrieveTransfer struct {
	Base
	TransferID string `json:"-"`
}

func (req *RetrieveTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transfer{}, &UpdateTransfer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type UpdateTransfer struct {
	Base
	TransferID string `json:"-"`
	Amount int `json:"amount"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (req *UpdateTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      PATCH,
		Path:        "/transfers/" + req.TransferID,
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transfer{}, &SearchTransfer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type SearchTransfer struct {
	Base
}

func (req *SearchTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/transfers/search",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transfer{}, &ListTransfers{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListTransfers struct {
	Base
}

func (req *ListTransfers) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/transfers",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transfer{}, &CreateTransfer{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type CreateTransfer struct {
	Base
	Amount int `json:"amount"`
	FailFast bool `json:"fail_fast"`
	Metadata map[string]interface{} `json:"metadata"`
	Recipient string `json:"recipient"`
}

func (req *CreateTransfer) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/transfers",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Schedule{}, &ListTransfersSchedules{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type ListTransfersSchedules struct {
	Base
}

func (req *ListTransfersSchedules) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      GET,
		Path:        "/transfers/schedules",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transfer{}, &MarkTransferAsPaid{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type MarkTransferAsPaid struct {
	Base
	TransferID string `json:"-"`
}

func (req *MarkTransferAsPaid) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/transfers/" + req.TransferID + "/mark_as_paid",
		ContentType: "application/json",
	}
}

// Example:
//
//	charge, update := &omise.Transfer{}, &MarkTransferAsSent{
//		ChargeID:    "chrg_456",
//		Description: "updated charge.",
//	}
//	if e := client.Do(charge, update); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("updated charge: %#v\n", charge)
//
type MarkTransferAsSent struct {
	Base
	TransferID string `json:"-"`
}

func (req *MarkTransferAsSent) Describe() *internal.Description {
	return &internal.Description{
		Endpoint:    internal.API,
		Method:      POST,
		Path:        "/transfers/" + req.TransferID + "/mark_as_sent",
		ContentType: "application/json",
	}
}

