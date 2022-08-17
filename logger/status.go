package logger

import (
	"fmt"
	"strings"

	"github.com/MCTosochu/MCSrvController/config"
)

func Status(up bool, config *config.ConfigStruct) {
	Text := strings.Join([]string{"ServerController Status", map[bool]string{true: "Up!", false: "Shutdown..."}[up]}, ": ")
	Color := map[bool]int32{true: 46417, false: 11730995}[up]

	fmt.Println(Text)
	SendDiscord(config, BaseDiscordEmbeds{
		Title: Text,
		Color: Color,
	})
}
