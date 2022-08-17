package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/MCTosochu/MCSrvController/config"
	"github.com/MCTosochu/MCSrvController/logger"
)

func Finisher(exitCode int, status *config.StatusSet) {
	logger.ControllerStatus(false, status.Config)
	if status.Srv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := status.Srv.Shutdown(ctx); err != nil {
			FatalError(err, status)
		}
	}
	os.Exit(exitCode)
}

func Fatal(text string, status *config.StatusSet) {
	fmt.Println(text)
	logger.SendDiscord(status.Config, logger.BaseDiscordEmbeds{
		Title: text,
		Color: 11730995,
	})
}

func FatalAndExit(text string, status *config.StatusSet) {
	Fatal(text, status)
	Finisher(1, status)
}

func FatalError(err error, status *config.StatusSet) {
	fmt.Println(err)
	logger.SendDiscord(status.Config, logger.BaseDiscordEmbeds{
		Title: fmt.Sprint(err),
		Color: 11730995,
	})
}

func FatalErrorAndExit(err error, status *config.StatusSet) {
	FatalError(err, status)
	Finisher(1, status)
}
