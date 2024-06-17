package main

import (
	"net/http"
)

func handleError(w http.ResponseWriter, status int, message string) {
	type errorMessage struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, 400, errorMessage{
		Error: message,
	})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	handleError(w, 200, "Something went wrong")
}
