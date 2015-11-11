package main

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/spf13/cobra"
)

var BalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Retrieve balance object.",
	RunE:  runBalance,
}

func runBalance(cmd *cobra.Command, args []string) error {
	return do(
		&omise.Balance{},
		&operations.RetrieveBalance{},
	)
}
