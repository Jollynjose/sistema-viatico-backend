package helpers

import (
	"log"
	"net/http"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
	log.Printf("Error: %s", message)
}
