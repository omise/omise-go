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
	switch len(args) {
	case 0:
		return ErrMissingArg("customer-id")
	}

	return do(
		&omise.CardList{},
		&operations.ListCards{
			CustomerID: args[0],
		},
	)
}

func runRetrieveCard(cmd *cobra.Command, args []string) error {
	// TODO: Ability to set customer-id in config.
	switch len(args) {
	case 0:
		return ErrMissingArg("customer-id")
	case 1:
		return ErrMissingArg("card-id")
	}

	return do(
		&omise.Card{},
		&operations.RetrieveCard{
			CustomerID: args[0],
			CardID:     args[1],
		},
	)
}
