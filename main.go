package main

import (
	"fmt"
	"os"

	"github.com/MCTosochu/MCSrvController/config"
)

func main() {
	config, configLoadError := config.Load()

	if configLoadError != nil {
		fmt.Println(configLoadError)
		os.Exit(1)
	}
}
