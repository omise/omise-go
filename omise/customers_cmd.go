package main

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/spf13/cobra"
)

var CustomersCmd = &cobra.Command{
	Use:   "customers",
	Short: "Manage customers.",
	RunE:  runListCustomers,
}

func init() {
	CustomersCmd.AddCommand(
		ListCustomersCmd,
		GetCustomerCmd,
		RemoveCustomerCmd,
	)
}

var ListCustomersCmd = &cobra.Command{
	Use:   "list",
	Short: "List customers.",
	RunE:  runListCustomers,
}

func runListCustomers(cmd *cobra.Command, args []string) error {
	return do(
		&omise.CustomerList{},
		&operations.ListCustomers{},
	)
}

var GetCustomerCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single customer.",
	RunE:  runGetCustomer,
}

func runGetCustomer(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "customer-id"); e != nil {
		return e
	}

	return do(
		&omise.Customer{},
		&operations.RetrieveCustomer{
			CustomerID: args[0],
		},
	)
}

// TODO: How to UpdateCustomer via CLI?
var RemoveCustomerCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a customer.",
	RunE:  runRemoveCustomer,
}

func runRemoveCustomer(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "customer-id"); e != nil {
		return e
	}

	return do(
		&omise.Deletion{},
		&operations.DestroyCustomer{
			CustomerID: args[0],
		},
	)
}
