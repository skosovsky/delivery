package app

import "delivery/internal/config"

type App struct{}

func New(cfg config.Config) *App {
	return &App{}
}
