package main

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/spf13/cobra"
)

var CardsCmd = &cobra.Command{
	Use:       "cards",
	Short:     "List customer's cards",
	ValidArgs: []string{"customer-id"},
	RunE:      runCards,
}

func init() {
	CardsCmd.AddCommand(RetrieveCardCmd)
}

var RetrieveCardCmd = &cobra.Command{
	Use:       "get",
	Short:     "Retrieve a single card object.",
	ValidArgs: []string{"card-id"},
	RunE:      runRetrieveCard,
}

func runCards(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "customer-id"); e != nil {
		return e
	}

	return do(
		&omise.CardList{},
		&operations.ListCards{
			CustomerID: args[0],
		},
	)
}

func runRetrieveCard(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "customer-id", "card-id"); e != nil {
		return e
	}

	return do(
		&omise.Card{},
		&operations.RetrieveCard{
			CustomerID: args[0],
			CardID:     args[1],
		},
	)
}
