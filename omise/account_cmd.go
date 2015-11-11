package main

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Retrieve account object.",
	RunE:  runAccount,
}

func runAccount(cmd *cobra.Command, args []string) error {
	return do(
		&omise.Account{},
		&operations.RetrieveAccount{},
	)
}
