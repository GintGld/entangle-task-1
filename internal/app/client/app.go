package app

import (
	client "task1/internal/client/ETH"
)

type App struct {
	C *client.Client
}

func New(url string) *App {
	return &App{
		C: client.New(url),
	}
}
