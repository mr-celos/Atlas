package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) { // handles the /health endpoint
	w.Header().Set("Content-Type", "application/json")
	data, err := healthJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e := ErrorResponse{Message: err.Error()}
		JSONValue, marshalErr := json.Marshal(e)
		if marshalErr != nil {
			fmt.Println("Error marshalling error response:", marshalErr)
			return
		}
		w.Write(JSONValue)
		return
	}
	w.Write(data)
}

func healthJSON() ([]byte, error) { // returns the health status as a JSON
	h := Health{Status: "OK", Version: "0.0.1"}

	JSONData, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}

	return JSONData, err
}
