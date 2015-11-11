package main

import (
	flag "github.com/spf13/pflag"
	"os"
)

var config = &struct {
	PKey string
	SKey string
}{}

func configure(flags *flag.FlagSet) {
	flags.StringVarP(&config.PKey, "pkey", "", "", "Omise API public key, defaults to $OMISE_PUBKEY.")
	flags.StringVarP(&config.SKey, "skey", "", "", "Omise API secret key, defaults to $OMISE_KEY.")
}

func envOverride() {
	// could do this with spf13/viper if we need more
	pkey, skey := os.Getenv("OMISE_PUBKEY"), os.Getenv("OMISE_KEY")
	if pkey != "" {
		config.PKey = pkey
	}
	if skey != "" {
		config.SKey = skey
	}
}
