package main

import (
	"net/http"

	"github.com/mrpbennett/docksmith/internal/logz"
)

var logger = logz.New()

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
	})

	logger.Info("webserver now running on 8080!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatalf("Fatal Error: %v", err)
	}
}
