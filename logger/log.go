package logger

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MCTosochu/MCSrvController/config"
)

func SendDiscord(config *config.ConfigStruct, data BaseDiscordEmbeds) {
	client := &http.Client{}

	dataJson := BaseDiscordText{
		Username: strings.Join([]string{config.Ident, " / Controller"}, ""),
		Embeds:   []BaseDiscordEmbeds{data},
	}

	payload, err := json.Marshal(dataJson)

	if err == nil {
		return
	}

	req, _ := http.NewRequest("POST", config.DiscordWebhookPath, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")

	res, _ := client.Do(req)
	res.Body.Close()
}

type BaseDiscordText struct {
	Username string              `json:"username"`
	Embeds   []BaseDiscordEmbeds `json:"embeds"`
}

type BaseDiscordEmbeds struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int32  `json:"color"`
}
