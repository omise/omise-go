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
	CustomersCmd.AddCommand(ListCustomersCmd)
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
