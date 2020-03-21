package main

import (
	_ "expvar"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/gops/agent"

	"github.com/bygui86/go-save-memory/http-server/config"
	"github.com/bygui86/go-save-memory/http-server/logging"
	"github.com/bygui86/go-save-memory/http-server/monitoring"
	"github.com/bygui86/go-save-memory/http-server/rest"
)

func main() {
	// from "github.com/pkg/profile"
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.GoroutineProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MutexProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.ThreadcreationProfile, profile.ProfilePath(".")).Stop()

	logging.Log.Info("Start http-server")

	cfg := config.LoadConfig()

	monitorServer := startMonitor()

	rest.RegisterCustomMetrics()

	restServer := startRestServer()

	startGopsAgent()

	startDebugServer()

	logging.Log.Info("http-server up&running")

	startSysCallChannel()

	shutdownAndWait(cfg, monitorServer, restServer)
}

func startMonitor() *monitoring.Server {
	logging.Log.Info("Start Monitoring server")

	server := monitoring.NewMonitorServer()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func startRestServer() *rest.Server {
	logging.Log.Info("Start REST server")

	server := rest.NewRestServer()
	logging.Log.Debug("HTTP server successfully created")

	server.Start()
	logging.Log.Debug("HTTP server successfully started")

	return server
}

func startGopsAgent() {
	logging.Log.Info("Start GoPS agent")

	err := agent.Listen(agent.Options{})
	if err != nil {
		logging.Log.Error("GoPS agent start failed:", err.Error())
		os.Exit(404)
	}
}

func startDebugServer() {
	logging.Log.Info("Start DEBUG server")
	go func() {
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			logging.Log.Error("Error starting REST server:", err)
		}
	}()
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(cfg *config.Config, monitorServer *monitoring.Server, restServer *rest.Server) {
	logging.Log.Warn("Termination signal received! Timeout", cfg.ShutdownTimeout)
	monitorServer.Shutdown(cfg.ShutdownTimeout)
	restServer.Shutdown(cfg.ShutdownTimeout)
	time.Sleep(time.Duration(cfg.ShutdownTimeout+1) * time.Second)
}
