package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bygui86/go-save-memory/http-server/logging"
)

func (s *Server) getUser(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("GET user request")

	setJsonContentType(writer)
	err := json.NewEncoder(writer).Encode(s.user)
	if err != nil {
		logging.SugaredLog.Errorf("Error on GET user request: %s", err.Error())
	}

	IncreaseGetRequests()
}

func (s *Server) postUser(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("POST user request")

	var responseMsg string
	decErr := json.NewDecoder(request.Body).Decode(s.user)
	if decErr != nil {
		responseMsg = fmt.Sprintf("POST user request failed: %s", decErr.Error())
	} else {
		responseMsg = "Set new user successful"
	}

	setJsonContentType(writer)
	setStatusAccepted(writer)
	encErr := json.NewEncoder(writer).Encode(&Result{Message: responseMsg})
	if encErr != nil {
		logging.SugaredLog.Errorf("Error on POST user request response encoding: %s", encErr.Error())
	}

	IncreasePostRequests()
}

func (s *Server) putUser(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("PUT user request")

	var responseMsg string
	var updatedUser User
	decErr := json.NewDecoder(request.Body).Decode(&updatedUser)
	if decErr != nil {
		responseMsg = fmt.Sprintf("PUT user request failed: %s", decErr.Error())
	} else {
		s.updateUser(&updatedUser)
		responseMsg = "Update user successful"
	}

	setJsonContentType(writer)
	setStatusAccepted(writer)
	encErr := json.NewEncoder(writer).Encode(&Result{Message: responseMsg})
	if encErr != nil {
		logging.SugaredLog.Errorf("Error on PUT user request response encoding: %s", encErr.Error())
	}

	IncreasePutRequests()
}

func (s *Server) deleteUser(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("DELETE user request")

	s.user = buildEmptyUser()

	setJsonContentType(writer)
	setStatusAccepted(writer)
	encErr := json.NewEncoder(writer).Encode(
		&Result{Message: "Delete user successful"},
	)
	if encErr != nil {
		logging.SugaredLog.Errorf("Error on DELETE user request response encoding: %s", encErr.Error())
	}

	IncreaseDeleteRequests()
}
