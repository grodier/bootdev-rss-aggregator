package responder

import (
	"log"
	"net/http"
)

func WithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Responding with 5xx error: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	WithJson(w, status, errorResponse{
		Error: msg,
	})
}
