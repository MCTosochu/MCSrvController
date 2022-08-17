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
	configData, configLoadError := config.Load()

	if configLoadError != nil {
		fmt.Println(configLoadError)
		os.Exit(1)
	}

	status := &(config.StatusSet{
		Config: configData,
		Srv:    nil,
	})

	logger.ControllerStatus(true, status.Config)

	time.Sleep(1 * time.Second)

	server.Listen(status)

	//Gracefully Shutdown
	func() {
		trap := make(chan os.Signal, 1)
		signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
		<-trap
		server.Finisher(0, status)
	}()
}
