package logger

import (
	"fmt"
	"strings"

	"github.com/MCTosochu/MCSrvController/config"
)

func status(target string, up bool, config *config.ConfigStruct) {
	Text := strings.Join([]string{target, " Status: ", map[bool]string{true: "Up!", false: "Shutdown..."}[up]}, "")
	Color := map[bool]int32{true: 46417, false: 11730995}[up]

	fmt.Println(Text)
	SendDiscord(config, BaseDiscordEmbeds{
		Title: Text,
		Color: Color,
	})
}

func ControllerStatus(up bool, config *config.ConfigStruct) {
	status("Controller", up, config)
}

func WebServerStatus(up bool, config *config.ConfigStruct) {
	status("WebServer", up, config)
}
