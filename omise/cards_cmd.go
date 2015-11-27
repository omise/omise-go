package main

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/spf13/cobra"
)

var CardsCmd = &cobra.Command{
	Use:   "cards",
	Short: "Manage cards.",
	RunE:  runListCards,
}

func init() {
	CardsCmd.AddCommand(
		ListCardsCmd,
		GetCardCmd,
		RemoveCardCmd,
	)
}

var ListCardsCmd = &cobra.Command{
	Use:   "list",
	Short: "List cards.",
	RunE:  runListCards,
}

func runListCards(cmd *cobra.Command, args []string) error {
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

var GetCardCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single card.",
	RunE:  runGetCard,
}

func runGetCard(cmd *cobra.Command, args []string) error {
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

var RemoveCardCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a card.",
	RunE:  runRemoveCard,
}

func runRemoveCard(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "customer-id", "card-id"); e != nil {
		return e
	}

	return do(
		&omise.Deletion{},
		&operations.DestroyCard{
			CustomerID: args[0],
			CardID:     args[1],
		},
	)
}
