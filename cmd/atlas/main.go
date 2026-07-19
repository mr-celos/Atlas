package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mr-celos/Atlas/internal/config"
	"github.com/mr-celos/Atlas/internal/logger"
)

func main() {
	logger.Init()
	cfg := config.Load()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthHandler(w, r, cfg.Version)
	})

	server := &http.Server{
		Addr: cfg.Port,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("ListenAndServe", "error", err)
			os.Exit(1)
		}
	}()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	<-channel

	slog.Info("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Graceful Shutdown", "error", err)
	}
}
