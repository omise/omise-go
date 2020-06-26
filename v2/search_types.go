package omise

import "context"

// SearchScope represents an enumeration of possible scopes that can be used with the
// Search API.
type SearchScope string

// SearchScope can be one of the following constants:
const (
    ChargeScope SearchScope = "charge"
    ChargeScheduleScope SearchScope = "charge_schedule"
    CustomerScope SearchScope = "customer"
    DisputeScope SearchScope = "dispute"
    EventScope SearchScope = "event"
    LinkScope SearchScope = "link"
    RecipientScope SearchScope = "recipient"
    RefundScope SearchScope = "refund"
    TransferScope SearchScope = "transfer"
    TransferScheduleScope SearchScope = "transfer_schedule"
    TransactionScope SearchScope = "transaction"
)

// ChargeSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type ChargeSearchResult struct {
	Search
	Data []*Charge `json:"data"`
}

// SearchCharges searchs Charge
//
// Example:
//
//	result, e := client.SearchCharge(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchCharges(ctx context.Context, params *SearchParams) (*ChargeSearchResult, error) {
	params.Scope = ChargeScope
	result := &ChargeSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// ChargeScheduleSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type ChargeScheduleSearchResult struct {
	Search
	Data []*ChargeSchedule `json:"data"`
}

// SearchChargeSchedules searchs ChargeSchedule
//
// Example:
//
//	result, e := client.SearchChargeSchedule(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchChargeSchedules(ctx context.Context, params *SearchParams) (*ChargeScheduleSearchResult, error) {
	params.Scope = ChargeScheduleScope
	result := &ChargeScheduleSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// CustomerSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type CustomerSearchResult struct {
	Search
	Data []*Customer `json:"data"`
}

// SearchCustomers searchs Customer
//
// Example:
//
//	result, e := client.SearchCustomer(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchCustomers(ctx context.Context, params *SearchParams) (*CustomerSearchResult, error) {
	params.Scope = CustomerScope
	result := &CustomerSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// DisputeSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type DisputeSearchResult struct {
	Search
	Data []*Dispute `json:"data"`
}

// SearchDisputes searchs Dispute
//
// Example:
//
//	result, e := client.SearchDispute(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchDisputes(ctx context.Context, params *SearchParams) (*DisputeSearchResult, error) {
	params.Scope = DisputeScope
	result := &DisputeSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// EventSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type EventSearchResult struct {
	Search
	Data []*Event `json:"data"`
}

// SearchEvents searchs Event
//
// Example:
//
//	result, e := client.SearchEvent(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchEvents(ctx context.Context, params *SearchParams) (*EventSearchResult, error) {
	params.Scope = EventScope
	result := &EventSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// LinkSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type LinkSearchResult struct {
	Search
	Data []*Link `json:"data"`
}

// SearchLinks searchs Link
//
// Example:
//
//	result, e := client.SearchLink(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchLinks(ctx context.Context, params *SearchParams) (*LinkSearchResult, error) {
	params.Scope = LinkScope
	result := &LinkSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RecipientSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type RecipientSearchResult struct {
	Search
	Data []*Recipient `json:"data"`
}

// SearchRecipients searchs Recipient
//
// Example:
//
//	result, e := client.SearchRecipient(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchRecipients(ctx context.Context, params *SearchParams) (*RecipientSearchResult, error) {
	params.Scope = RecipientScope
	result := &RecipientSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// RefundSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type RefundSearchResult struct {
	Search
	Data []*Refund `json:"data"`
}

// SearchRefunds searchs Refund
//
// Example:
//
//	result, e := client.SearchRefund(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchRefunds(ctx context.Context, params *SearchParams) (*RefundSearchResult, error) {
	params.Scope = RefundScope
	result := &RefundSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// TransferSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type TransferSearchResult struct {
	Search
	Data []*Transfer `json:"data"`
}

// SearchTransfers searchs Transfer
//
// Example:
//
//	result, e := client.SearchTransfer(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchTransfers(ctx context.Context, params *SearchParams) (*TransferSearchResult, error) {
	params.Scope = TransferScope
	result := &TransferSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// TransferScheduleSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type TransferScheduleSearchResult struct {
	Search
	Data []*TransferSchedule `json:"data"`
}

// SearchTransferSchedules searchs TransferSchedule
//
// Example:
//
//	result, e := client.SearchTransferSchedule(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchTransferSchedules(ctx context.Context, params *SearchParams) (*TransferScheduleSearchResult, error) {
	params.Scope = TransferScheduleScope
	result := &TransferScheduleSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

// TransactionSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type TransactionSearchResult struct {
	Search
	Data []*Transaction `json:"data"`
}

// SearchTransactions searchs Transaction
//
// Example:
//
//	result, e := client.SearchTransaction(search)
//
//	fmt.Printf("search: %#v\n", result)
//
func (s *SearchService) SearchTransactions(ctx context.Context, params *SearchParams) (*TransactionSearchResult, error) {
	params.Scope = TransactionScope
	result := &TransactionSearchResult{}
	err := s.Do(ctx, result, params)

	return result, err
}

