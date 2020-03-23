package rest

import (
	"net/http"

	"github.com/bygui86/go-save-memory/http-server/logging"
)

const (
	contentTypeHeaderKey   = "Content-Type"
	contentTypeHeaderValue = "application/json"
)

// setJsonContentType - Set response header "Content-Type" to "application/json"
func setJsonContentType(writer http.ResponseWriter) {
	writer.Header().Set(contentTypeHeaderKey, contentTypeHeaderValue)
}

// setStatusAccepted - Set response status code to 202-Accepted
func setStatusAccepted(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusAccepted)
}

// buildEmptyUser - Build an User with default values
func buildEmptyUser() *User {
	return &User{
		Name:       "",
		Surname:    "",
		Age:        0,
		BestMovies: make(map[string]string, 1),
		Hobbies:    make([]string, 1),
	}
}

// updateUser - Update the existing User with new values
func (s *Server) updateUser(updatedUser *User) {
	if updatedUser.Name != "" {
		s.user.Name = updatedUser.Name
	} else {
		logging.SugaredLog.Warnf("Updated user name empty, skipping")
	}

	if updatedUser.Surname != "" {
		s.user.Surname = updatedUser.Surname
	} else {
		logging.SugaredLog.Warnf("Updated user surname empty, skipping")
	}

	if updatedUser.Age > 0 {
		s.user.Age = updatedUser.Age
	} else {
		logging.SugaredLog.Warnf("Updated user age negative or zero, skipping")
	}

	if len(updatedUser.BestMovies) > 0 {
		s.user.BestMovies = updatedUser.BestMovies
	} else {
		logging.SugaredLog.Warnf("Updated user best movies empty, skipping")
	}

	if len(updatedUser.Hobbies) > 0 {
		s.user.Hobbies = updatedUser.Hobbies
	} else {
		logging.SugaredLog.Warnf("Updated user hobbies empty, skipping")
	}
}
