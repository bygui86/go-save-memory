package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/bygui86/go-save-memory/http-server/logging"

	"github.com/gorilla/mux"
)

// NewMonitorServer - Create new Monitoring server
func NewMonitorServer() *Server {
	logging.Log.Debug("Create new Monitoring server")

	cfg := loadConfig()

	server := &Server{
		Config: cfg,
	}

	server.setupRouter()
	server.setupHTTPServer()
	return server
}

// setupRouter - Create new Gorilla mux Router
func (s *Server) setupRouter() {
	logging.Log.Debug("Create new Router")

	s.Router = mux.NewRouter().StrictSlash(true)
	s.Router.Handle("/metrics", promhttp.Handler())
}

// newHTTPServer - Create new HTTP server
func (s *Server) setupHTTPServer() {
	logging.SugaredLog.Debugf("Create new HTTP server on port %d", s.Config.MonitorPort)

	if s.Config != nil {
		s.HTTPServer = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", s.Config.MonitorHost, s.Config.MonitorPort),
			Handler: s.Router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}
		return
	}

	logging.Log.Error("Monitoring server creation failed: Monitoring server configurations not initialized")
}

// Start - Start Monitoring server
func (s *Server) Start() {
	logging.Log.Info("Start Monitoring server")

	if s.HTTPServer != nil && !s.Running {
		go func() {
			err := s.HTTPServer.ListenAndServe()
			if err != nil {
				logging.SugaredLog.Errorf("Error starting Monitoring server: %s", err.Error())
			}
		}()
		s.Running = true
		logging.SugaredLog.Infof("Monitoring server listening on port %d", s.Config.MonitorPort)
		return
	}

	logging.Log.Error("Monitoring server start failed: HTTP server not initialized or Monitoring server already running")
}

// Shutdown - Shutdown Monitoring server
func (s *Server) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown Monitoring server, timeout %d", timeout)

	if s.HTTPServer != nil && s.Running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.HTTPServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("Error shutting down Monitoring server: %s", err)
		}
		s.Running = false
		return
	}

	logging.Log.Error("Monitoring server shutdown failed: HTTP server not initialized or Monitoring server not running")
}
