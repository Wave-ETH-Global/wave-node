package router

import (
	"github.com/Wave-ETH-Global/wave-node/app"
	"github.com/labstack/echo/v4"
)

func NewRouter(app *app.App) (*echo.Echo, error) {
	router := echo.New()
	router.HideBanner = true

	return router, nil
}
