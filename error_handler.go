package main

import (
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, status int, message string) {
	if status > 399 {
		log.Println("Something went wrong")
		return
	}

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
