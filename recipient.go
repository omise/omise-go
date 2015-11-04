package omise

type Recipient struct {
	Base
	Verified    bool          `json:"verified"`
	Active      bool          `json:"active"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Description *string       `json:"description"`
	Type        RecipientType `json:"type"`
	TaxID       *string       `json:"tax_id"`
	BankAccount *BankAccount  `json:"bank_account"`
	FailureCode *string       `json:"failure_code"`
}

type RecipientList struct {
	List
	Data []*Recipient `json:"data"`
}

func (list *RecipientList) Find(id string) *Recipient {
	for _, recipient := range list.Data {
		if recipient.ID == id {
			return recipient
		}
	}

	return nil
}
