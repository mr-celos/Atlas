package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mr-celos/Atlas/internal/config"
	"github.com/mr-celos/Atlas/internal/database"
	"github.com/mr-celos/Atlas/internal/logger"
)

func main() {
	logger.Init()
	cfg := config.Load()

	pool, poolErr := database.Connect(cfg)
	if poolErr != nil {
		slog.Error("Connect to Database", "error", poolErr)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthHandler(w, r, cfg.Version, pool)
	})

	server := &http.Server{
		Addr: cfg.Port,
	}

	go func() { // starts the server in a separate goroutine so that it doesn't block the main thread
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("ListenAndServe", "error", err)
			os.Exit(1)
		}
	}()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	<-channel // blocks the main thread until an interrupt signal is received

	slog.Info("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	defer pool.Close()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Graceful Shutdown", "error", err)
	}
}
