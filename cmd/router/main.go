package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"task1/internal/app"
	"task1/internal/config"
)

func main() {
	// Setup config.
	cfg := config.MustLoad()

	// Setup logger.
	log := setupLogger()

	log.Info("starting server", slog.Int("port", cfg.Port))
	log.Debug("debug messages are enabled")

	// Initialize app.
	httpApplication := app.New(
		log,
		cfg.Url,
		cfg.Port,
	)

	// Run server.
	go httpApplication.Router.Run()

	// Graceful shutdown.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	<-stop

	// Stop application.
	httpApplication.Router.Stop()
	log.Info("Gracefully stopped")
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)
}
