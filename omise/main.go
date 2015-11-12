package main

import (
	"encoding/json"
	"log"
	"os"

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
	)
}

func main() {
	configure(OmiseCmd.PersistentFlags())
	if e := OmiseCmd.Execute(); e != nil {
		log.Fatalln(e)
	}
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

func checkArgs(args []string, argnames ...string) error {
	if len(args) == len(argnames) {
		return nil
	}

	return ErrMissingArg(argnames[len(args)])
}

func output(obj interface{}) error {
	bytes, e := json.MarshalIndent(obj, "", "  ")
	if e != nil {
		return e
	}

	bytes = append(bytes, []byte("\n")...)
	_, e = os.Stdout.Write(bytes)
	return e
}
