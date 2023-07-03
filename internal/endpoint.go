package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Endpoint string

type EnvConfig struct {
	ApiUrl   Endpoint
	VaultUrl Endpoint
}

func GetEnv() *EnvConfig {
	var defaultUrls = &EnvConfig{
		ApiUrl:   Endpoint("https://api.omise.co"),
		VaultUrl: Endpoint("https://vault.omise.co"),
	}
	var err = godotenv.Load()

	if err != nil {
		return defaultUrls
	}

	// Retrieve ApiUrl value from env variable
	var apiUrlEnv, exists = os.LookupEnv("OMISE_GO_API_URL")

	// If API URL is not present then no need to look in to Vault URL
	// because we require both URLs to set up new URLs.
	if !exists {
		return defaultUrls
	}

	var vaultUrlEnv string
	// Retrieve OMISE_GO_VAULT_URL value from env variable
	vaultUrlEnv, exists = os.LookupEnv("OMISE_GO_VAULT_URL")

	// If Vault URL is not present then we return default URLs
	// because we require both URLs to set up new URLs.
	if !exists {
		return defaultUrls
	}

	return &EnvConfig{
		ApiUrl:   Endpoint(apiUrlEnv),
		VaultUrl: Endpoint(vaultUrlEnv),
	}
}
