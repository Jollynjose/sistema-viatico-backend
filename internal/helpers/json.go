package helpers

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return err
	}
	return nil
}

func ParseJSON(w http.ResponseWriter, data interface{}) error {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return err
	}

	return nil
}
