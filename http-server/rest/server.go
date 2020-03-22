package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bygui86/go-save-memory/http-server/logging"

	"github.com/gorilla/mux"
)

// NewRestServer - Create new REST server
func NewRestServer() *Server {
	logging.Log.Debug("Create new REST server")

	cfg := loadConfig()

	server := &Server{
		Config: cfg,
		user:   buildEmptyUser(),
	}

	server.setupRouter()
	server.setupHTTPServer()
	return server
}

// setupRouter - Create new Gorilla mux Router
func (s *Server) setupRouter() {
	logging.Log.Debug("Create new Router")

	s.Router = mux.NewRouter().StrictSlash(true)

	s.Router.HandleFunc("/user", s.getUser).Methods(http.MethodGet)
	s.Router.HandleFunc("/user", s.postUser).Methods(http.MethodPost)
	s.Router.HandleFunc("/user", s.putUser).Methods(http.MethodPut)
	s.Router.HandleFunc("/user", s.deleteUser).Methods(http.MethodDelete)
}

// setupHTTPServer - Create new HTTP server
func (s *Server) setupHTTPServer() {
	logging.SugaredLog.Debugf("Create new HTTP server on port %d", s.Config.RestPort)

	if s.Config != nil {
		s.HTTPServer = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", s.Config.RestHost, s.Config.RestPort),
			Handler: s.Router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}
		return
	}

	logging.Log.Error("HTTP server creation failed: REST server configurations not initialized")
}

// Start - Start REST server
func (s *Server) Start() {
	logging.Log.Info("Start REST server")

	if s.HTTPServer != nil && !s.Running {
		go func() {
			err := s.HTTPServer.ListenAndServe()
			if err != nil {
				logging.SugaredLog.Errorf("Error starting REST server: %s", err.Error())
			}
		}()
		s.Running = true
		logging.SugaredLog.Infof("REST server listening on port %d", s.Config.RestPort)
		return
	}

	logging.Log.Error("REST server start failed: HTTP server not initialized or HTTP server already running")
}

// Shutdown - Shutdown REST server
func (s *Server) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown REST server, timeout %d", timeout)

	if s.HTTPServer != nil && s.Running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.HTTPServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("Error shutting down REST server: %s", err.Error())
		}
		s.Running = false
		return
	}

	logging.Log.Error("REST server shutdown failed: HTTP server not initialized or HTTP server not running")
}
