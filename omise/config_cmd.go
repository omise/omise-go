package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage customer's cards.",
	RunE:  runConfig,
}

func init() {
	ConfigCmd.AddCommand(ConfigGetCmd, ConfigSetCmd)
}

var ConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a configuration value.",
	RunE:  runConfigGet,
}

var ConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a configuration value and saves it to ~/.omise.",
	RunE:  runConfigSet,
}

func runConfig(cmd *cobra.Command, args []string) error {
	configs := map[string]string{}
	for _, key := range config.viper.AllKeys() {
		configs[key] = fmt.Sprint(config.viper.Get(key))
	}

	return output(configs)
}

func runConfigGet(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "config-key"); e != nil {
		return e
	}

	value := fmt.Sprint(config.viper.Get(args[0])) + "\n"
	_, e := os.Stdout.Write([]byte(value))
	return e
}

func runConfigSet(cmd *cobra.Command, args []string) error {
	if e := checkArgs(args, "config-key", "config-value"); e != nil {
		return e
	}

	key, value := args[0], args[1]
	return editConfigYAML(func(configs map[string]string) error {
		configs[key] = value
		return nil
	})
}

func editConfigYAML(action func(map[string]string) error) error {
	configPath := os.Getenv("HOME") + "/.omise"
	bytes, e := ioutil.ReadFile(configPath)
	if e != nil {
		return e
	}

	configs := map[string]string{}
	if e := json.Unmarshal(bytes, &configs); e != nil {
		return e
	}
	if e := action(configs); e != nil {
		return e
	}

	if bytes, e = json.MarshalIndent(configs, "", "  "); e != nil {
		return e
	}

	return ioutil.WriteFile(configPath, bytes, 0644)
}
