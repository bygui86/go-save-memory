package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
	Running    bool
	user       *User
}

// WARN: remember to expose all fields externally, otherwise the json marshaler won't be able to find and use them
type User struct {
	Name       string            `json:"name"`
	Surname    string            `json:"surname"`
	Age        int               `json:"age"`
	BestMovies map[string]string `json:"bestMovies"` // key: genre - value: title
	Hobbies    []string          `json:"hobbies"`
}

type Result struct {
	Message string `json:"Message"`
}
