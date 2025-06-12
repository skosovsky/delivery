package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"delivery/internal/app"
	"delivery/internal/config"
)

func main() {
	configs := config.New()
	application := app.New(configs)

	startWebServer(application, configs.HTTPPort)
}

func startWebServer(_ *app.App, port string) {
	http.HandleFunc(http.MethodGet+" /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Healthy"))
		if err != nil {
			slog.Warn("Error writing response", "error", err)
		}
	})

	address := fmt.Sprintf(":%s", port)
	slog.Info("Starting server", "address", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		slog.Error("Server failed", "error", err)
	}
}
