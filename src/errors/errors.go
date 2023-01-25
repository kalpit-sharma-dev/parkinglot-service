package errors

import (
	"encoding/json"
	"net/http"
)

type ParkingCustomError struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func HandleError(w http.ResponseWriter, errResp error) {
	respBytes, err := json.Marshal(errResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(respBytes)
}
