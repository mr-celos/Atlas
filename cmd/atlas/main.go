package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/mr-celos/Atlas/internal/config"
	"github.com/mr-celos/Atlas/internal/logger"
)

func main() {
	logger.Init()
	cfg := config.Load()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthHandler(w, r, cfg.Version)
	})

	err := http.ListenAndServe(cfg.Port, nil)
	if err != nil {
		slog.Error("ListenAndServe", "error", err)
		os.Exit(1)
	}
}
