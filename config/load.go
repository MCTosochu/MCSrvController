package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfigStruct struct {
	Ident              string
	DiscordWebhookPath string
	Port               uint32
}

func Load() (config *ConfigStruct, err error) {

	f, fileErr := ioutil.ReadFile("requirements.json")

	if fileErr != nil {
		return nil, fmt.Errorf("requirements.json not found.")
	}

	if jsonErr := json.NewDecoder(bytes.NewReader(f)).Decode(&config); jsonErr != nil {
		return nil, fmt.Errorf("requirements.json is wrong.")
	}

	return
}
