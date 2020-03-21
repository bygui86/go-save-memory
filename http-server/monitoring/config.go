package monitoring

import (
	"github.com/bygui86/go-save-memory/http-server/logging"
	"github.com/bygui86/go-save-memory/http-server/utils"
)

const (
	monHostEnvVar = "MONITOR_HOST"
	monPortEnvVar = "MONITOR_PORT"

	monHostDefault = "localhost"
	monPortDefault = 9090
)

type Config struct {
	RestHost string
	RestPort int
}

func loadConfig() *Config {
	logging.Log.Debug("Load Monitoring configurations")
	return &Config{
		RestHost: utils.GetStringEnv(monHostEnvVar, monHostDefault),
		RestPort: utils.GetIntEnv(monPortEnvVar, monPortDefault),
	}
}
