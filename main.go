package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MCTosochu/MCSrvController/config"
	"github.com/MCTosochu/MCSrvController/logger"
	"github.com/MCTosochu/MCSrvController/server"
)

func main() {
	config, configLoadError := config.Load()

	if configLoadError != nil {
		fmt.Println(configLoadError)
		os.Exit(1)
	}

	//Gracefully Shutdown
	go func() {
		trap := make(chan os.Signal, 1)
		signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
		<-trap
		server.Finisher(0, config)
	}()

	logger.Status(true, config)

	time.Sleep(1 * time.Second)

	server.Listen(config)

}
