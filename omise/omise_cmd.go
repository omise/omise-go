package main

import (
	"log"

	"github.com/spf13/cobra"
)

var OmiseCmd = &cobra.Command{
	Use:   "omise",
	Short: "omise cli command lets you perform omise API calls from the command line interactively.",
	RunE:  runOmise,

	PersistentPreRunE:  preRunOmise,
	PersistentPostRunE: postRunOmise,
}

func init() {
	OmiseCmd.AddCommand(
		ConfigCmd,
		AccountCmd,
		BalanceCmd,
		CardsCmd,
		CustomersCmd,
		ChargesCmd,
	)
}

func preRunOmise(cmd *cobra.Command, args []string) error {
	return bindViper(cmd)
}

func runOmise(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func postRunOmise(cmd *cobra.Command, args []string) error {
	log.Println("done.")
	return nil
}
