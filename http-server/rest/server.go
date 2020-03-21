package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bygui86/go-save-memory/http-server/logging"

	"github.com/gorilla/mux"
)

type Server struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
	Running    bool
	user       *User
}

// NewRestServer - Create new REST server
func NewRestServer() *Server {
	logging.Log.Info("Create new REST server")

	// create config
	cfg := loadConfig()

	// create router
	router := newRouter()

	// create http server
	httpServer := newHTTPServer(cfg.RestHost, cfg.RestPort, router)

	return &Server{
		Config:     cfg,
		Router:     router,
		HTTPServer: httpServer,
	}
}

// newRouter -
func newRouter() *mux.Router {
	logging.Log.Debug("Create new Router")

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/user", getUser).Methods(http.MethodGet)
	router.HandleFunc("/user", setUser).Methods(http.MethodPost)
	return router
}

// newHttpServer -
func newHTTPServer(host string, port int, router *mux.Router) *http.Server {
	logging.Log.Debug("Create new HTTP server on port", port)

	return &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

// Start - Start REST server
func (s *Server) Start() {
	logging.Log.Info("Start REST server")

	go func() {
		err := s.HTTPServer.ListenAndServe()
		if err != nil {
			logging.Log.Error("Error starting REST server:", err)
		}
	}()

	s.Running = true
	logging.Log.Info("REST server listening on port", s.Config.RestPort)
}

// Shutdown - Shutdown REST server
func (s *Server) Shutdown(timeout int) {
	logging.Log.Warn("Shutdown REST server, timeout", timeout)
	if s.HTTPServer != nil && s.Running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.HTTPServer.Shutdown(ctx)
		if err != nil {
			logging.Log.Error("Error shutting down REST server:", err)
		}
		s.Running = false
	}
}
