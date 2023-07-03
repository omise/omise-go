package internal

import (
	"os"

	"github.com/godotenv"
)

type EnvConfig struct {
	AppUrl   string
	VaultUrl string
}

func GetEnv() *EnvConfig {
	var defaultUrls = &EnvConfig{
		AppUrl:   "https://api.omise.co",
		VaultUrl: "https://vault.omise.co",
	}
	var err = godotenv.Load()

	if err != nil {
		return defaultUrls
	}

	// Retrieve APP_URL value from env variable
	var appUrlEnv, exists = os.LookupEnv("APP_URL")

	if !exists {
		return defaultUrls
	}

	var vaultUrlEnv string
	// Retrieve VAULT_URL value from env variable
	vaultUrlEnv, exists = os.LookupEnv("VAULT_URL")

	if !exists {
		return defaultUrls
	}

	return &EnvConfig{
		AppUrl:   appUrlEnv,
		VaultUrl: vaultUrlEnv,
	}
}
