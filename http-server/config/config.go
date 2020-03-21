package config

import (
	"github.com/bygui86/go-save-memory/http-server/logging"
	"github.com/bygui86/go-save-memory/http-server/utils"
)

const (
	shutdownTimeoutEnvVar  = "SHUTDOWN_TIMEOUT"
	shutdownTimeoutDefault = 10
)

type Config struct {
	ShutdownTimeout int
}

func LoadConfig() *Config {
	logging.Log.Debug("Load HTTP server general configurations")
	return &Config{
		ShutdownTimeout: utils.GetIntEnv(shutdownTimeoutEnvVar, shutdownTimeoutDefault),
	}
}
