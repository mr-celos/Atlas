package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Health struct {
	Status         string `json:"status"`
	Version        string `json:"version"`
	DatabaseStatus string `json:"database_status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request, version string, pool *pgxpool.Pool) { // handles the /health endpoint
	w.Header().Set("Content-Type", "application/json")
	data, err := healthJSON(version, pool)
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

func healthJSON(version string, pool *pgxpool.Pool) ([]byte, error) { // returns the health status as a JSON
	dbStatus := "FAIL"

	if pool != nil {
		dbStatus = "OK"
		pingErr := pool.Ping(context.Background()) //checks if the database connection is alive every time the /health endpoint is called

		if pingErr != nil {
			dbStatus = "FAIL"
		}
	}

	h := Health{Status: "OK", Version: version, DatabaseStatus: dbStatus}

	JSONData, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}

	return JSONData, err
}
