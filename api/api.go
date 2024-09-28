package api

import (
	"encoding/json"
	"net/http"
)

type UserNotesParams struct {
	Username string
}

type UserNotesResponse struct {
	Code  int
	Notes []string
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		writeError(w, err.Error(), http.StatusInternalServerError)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)
	}
)
