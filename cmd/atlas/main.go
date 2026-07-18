package main

import (
	"log"
	"net/http"
	"github.com/mr-celos/Atlas/internal/config"
)

func main() {
	cfg := config.Load()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthHandler(w, r, cfg.Version)
	})

	err := http.ListenAndServe(cfg.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
