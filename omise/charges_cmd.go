package main

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/spf13/cobra"
)

var ChargesCmd = &cobra.Command{
	Use:   "charges",
	Short: "Manage charges.",
	RunE:  runListCharges,
}

func init() {
	ChargesCmd.AddCommand(
		ListChargesCmd,
		GetChargeCmd,
	)
}

var ListChargesCmd = &cobra.Command{
	Use:   "list",
	Short: "List charges.",
	RunE:  runListCharges,
}

func runListCharges(cmd *cobra.Command, args []string) error {
	return do(
		&omise.ChargeList{},
		&operations.ListCharges{},
	)
}

// TODO: Create and update.

var GetChargeCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single charge.",
	RunE:  runGetCharge,
}

func runGetCharge(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "charge-id"); e != nil {
		return e
	}

	return do(
		&omise.Charge{},
		&operations.RetrieveCharge{
			ChargeID: args[0],
		},
	)
}
