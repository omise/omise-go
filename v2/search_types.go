package omise

// SearchScope represents an enumeration of possible scopes that can be used with the
// Search API.
type SearchScope string

// SearchScope can be one of the following constants:
const (
    ChargeScope SearchScope = "charge"
    ChargeSchedulingScope SearchScope = "charge_schedule"
    CustomerScope SearchScope = "customer"
    DisputeScope SearchScope = "dispute"
    EventScope SearchScope = "event"
    LinkScope SearchScope = "link"
    RecipientScope SearchScope = "recipient"
    RefundScope SearchScope = "refund"
    TransferScope SearchScope = "transfer"
    TransferSchedulingScope SearchScope = "transfer_schedule"
    TransactionScope SearchScope = "transaction"
)

// ChargeSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type ChargeSearchResult struct {
	Search
	Data []*Charge `json:"data"`
}

// ChargeSchedulingSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type ChargeSchedulingSearchResult struct {
	Search
	Data []*ChargeScheduling `json:"data"`
}

// CustomerSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type CustomerSearchResult struct {
	Search
	Data []*Customer `json:"data"`
}

// DisputeSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type DisputeSearchResult struct {
	Search
	Data []*Dispute `json:"data"`
}

// EventSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type EventSearchResult struct {
	Search
	Data []*Event `json:"data"`
}

// LinkSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type LinkSearchResult struct {
	Search
	Data []*Link `json:"data"`
}

// RecipientSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type RecipientSearchResult struct {
	Search
	Data []*Recipient `json:"data"`
}

// RefundSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type RefundSearchResult struct {
	Search
	Data []*Refund `json:"data"`
}

// TransferSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type TransferSearchResult struct {
	Search
	Data []*Transfer `json:"data"`
}

// TransferSchedulingSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type TransferSchedulingSearchResult struct {
	Search
	Data []*TransferScheduling `json:"data"`
}

// TransactionSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type TransactionSearchResult struct {
	Search
	Data []*Transaction `json:"data"`
}

