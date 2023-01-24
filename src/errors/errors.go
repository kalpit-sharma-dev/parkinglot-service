package errors

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, errResp error) {
	respBytes, err := json.Marshal(errResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	w.Write(respBytes)
}
