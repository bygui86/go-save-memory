package monitoring

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
	Running    bool
}
