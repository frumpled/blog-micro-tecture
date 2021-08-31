package controller

import (
	"encoding/json"
	"net/http"
)

func encodeResponse(w http.ResponseWriter, responseCode int, responseBody interface{}) {
	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	w.Write(responseJSON)
}

func decodeRequest(r *http.Request) (interface{}, error) {
	var request struct{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
