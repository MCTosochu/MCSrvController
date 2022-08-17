package config

import (
	"fmt"
	"os"
)

type SrvCtlConfig struct {
	RepositoryPath     string
	DiscordWebhookPath string
}

func Load() (*SrvCtlConfig, error) {

	f, err := os.Open("requirements.json")

	if err != nil {
		return nil, fmt.Errorf("requirements.json not found.")
	}

	defer f.Close()

	return nil, nil
}
