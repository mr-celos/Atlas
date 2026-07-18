package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}
