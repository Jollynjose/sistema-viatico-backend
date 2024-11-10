package helpers

import "net/http"

func ResponseHandler[T any](w http.ResponseWriter, statusCode int, response T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := ParseJSON(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
