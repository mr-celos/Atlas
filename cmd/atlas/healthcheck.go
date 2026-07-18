package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request, version string) { // handles the /health endpoint
	w.Header().Set("Content-Type", "application/json")
	data, err := healthJSON(version)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e := ErrorResponse{Message: err.Error()}
		JSONValue, marshalErr := json.Marshal(e)
		if marshalErr != nil {
			slog.Error("Failed to marshal error response", "error", marshalErr)
			return
		}
		w.Write(JSONValue)
		return
	}
	w.Write(data)
}

func healthJSON(version string) ([]byte, error) { // returns the health status as a JSON
	h := Health{Status: "OK", Version: version}

	JSONData, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}

	return JSONData, err
}
