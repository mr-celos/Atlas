package logger

import (
	"log/slog"
	"os"
)

func Init() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	log := slog.New(handler)

	slog.SetDefault(log)
}
