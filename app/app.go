package app

import (
	"context"
	"fmt"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/google/logger"
	"github.com/labstack/echo/v4"
)

type App struct {
	Config *config.Config
}

func NewApp(cfg *config.Config) (*App, error) {
	app := &App{}

	app.Config = cfg

	return app, nil
}

func (app *App) Run(ctx context.Context, router *echo.Echo) error {
	port := app.Config.HTTP.Port
	addr := app.Config.HTTP.Address
	logger.Infof("HTTP server starts listening at %s:%d", addr, port)
	go logger.Fatal(router.Start(fmt.Sprintf("%s:%d", addr, port)))
	return nil
}
