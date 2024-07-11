package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParsePayload(body io.ReadCloser, payload any) error {
	if body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(body).Decode(payload)
}

func Respond(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}

func RespondError(w http.ResponseWriter, status int, err error) {
	Respond(w, status, map[string]string{"error": err.Error()})
}
