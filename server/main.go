package server

import (
	"fmt"
	"net/http"

	"github.com/MCTosochu/MCSrvController/config"
	"github.com/MCTosochu/MCSrvController/logger"
)

var globalStatus *config.StatusSet

func Listen(status *config.StatusSet) {
	globalStatus = status
	status.Srv = &http.Server{Addr: ":" + globalStatus.Config.Port}

	go func() {
		logger.WebServerStatus(true, globalStatus.Config)
		err := status.Srv.ListenAndServe()
		if err != http.ErrServerClosed {
			FatalErrorAndExit(
				fmt.Errorf("WebServer Listen Error: %w", err),
				globalStatus,
			)
		}
	}()

}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready")
}
