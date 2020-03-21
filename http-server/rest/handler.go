package rest

import (
	"encoding/json"
	"net/http"

	"github.com/bygui86/go-save-memory/http-server/logging"

	"github.com/gorilla/mux"
)

const (
	defaultMsg = "Hello world!"
)

func getUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	json.Marshal()

	if len(msg) == 0 {
		logging.Log.Infof("Echo of default msg '%s'", defaultMsg)
		w.Write([]byte(defaultMsg))
	} else {
		logging.Log.Infof("Echo of msg '%s'", msg)
		w.Write([]byte(msg))
	}

	IncreaseOpsProcessed()
}
