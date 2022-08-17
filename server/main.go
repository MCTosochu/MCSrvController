package server

import (
	"fmt"
	"net/http"

	"github.com/MCTosochu/MCSrvController/config"
)

func Listen(config *config.ConfigStruct) {
	http.HandleFunc("/", version)
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		FatalAndExit(
			fmt.Sprint(fmt.Errorf("WebServer Listen Error: %w", err)),
			config,
		)
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready")
}
