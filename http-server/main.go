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
	logging.Log.Debug("Start Monitoring server")

	server := monitoring.NewMonitorServer()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func startRestServer() *rest.Server {
	logging.Log.Debug("Start REST server")

	server := rest.NewRestServer()
	logging.Log.Debug("REST server successfully created")

	server.Start()
	logging.Log.Debug("REST server successfully started")

	return server
}

func startGopsAgent() {
	logging.Log.Info("Start GoPS agent")

	err := agent.Listen(agent.Options{})
	if err != nil {
		logging.SugaredLog.Errorf("GoPS agent start failed: %s", err.Error())
		os.Exit(500)
	}

	logging.Log.Info("GoPS agent ready")
}

func startDebugServer() {
	logging.Log.Info("Start DEBUG server")

	go func() {
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			logging.SugaredLog.Errorf("Error starting REST server: %s", err.Error())
		}
	}()

	logging.Log.Info("DEBUG server listening on port 6060")
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(cfg *config.Config, monitorServer *monitoring.Server, restServer *rest.Server) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", cfg.ShutdownTimeout)
	monitorServer.Shutdown(cfg.ShutdownTimeout)
	restServer.Shutdown(cfg.ShutdownTimeout)
	time.Sleep(time.Duration(cfg.ShutdownTimeout+1) * time.Second)
}
