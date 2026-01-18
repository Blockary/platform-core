package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return err
	}
	defer r.Body.Close()
	return nil
}

func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if jsonBytes, err := json.Marshal(v); err == nil {
		w.Write(jsonBytes)
	} else {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
