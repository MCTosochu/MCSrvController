package config

import (
	"net/http"
)

type StatusSet struct {
	Config *ConfigStruct
	Srv    *http.Server
}
