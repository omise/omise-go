package main

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Prints configuration variables.",
	RunE:  runConfig,
}

func runConfig(cmd *cobra.Command, args []string) error {
	return output(config)
}
