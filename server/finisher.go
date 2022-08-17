package server

import (
	"fmt"
	"os"

	"github.com/MCTosochu/MCSrvController/config"
	"github.com/MCTosochu/MCSrvController/logger"
)

func Finisher(exitCode int, config *config.ConfigStruct) {
	logger.ControllerStatus(false, config)
	os.Exit(exitCode)
}

func FatalAndExit(text string, config *config.ConfigStruct) {
	fmt.Println(text)
	logger.SendDiscord(config, logger.BaseDiscordEmbeds{
		Title: text,
		Color: 11730995,
	})
	Finisher(1, config)
}
