package main

import (
	"fmt"
	"os"

	"github.com/MCTosochu/MCSrvController/config"
	"github.com/MCTosochu/MCSrvController/logger"
)

func main() {
	config, configLoadError := config.Load()

	if configLoadError != nil {
		fmt.Println(configLoadError)
		os.Exit(1)
	}

	logger.Status(true, config)
}
