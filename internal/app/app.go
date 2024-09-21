package app

import (
	"log/slog"
	client "task1/internal/app/client"
	router "task1/internal/app/router"
)

type App struct {
	log    *slog.Logger
	Router *router.App
	client *client.App
}

func New(
	log *slog.Logger,
	url string,
	port int,
) *App {
	client := client.New(url)
	router := router.New(
		log,
		port,
		client.C,
	)

	return &App{
		log:    log,
		Router: router,
		client: client,
	}
}
