package router

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/labstack/echo/v4"
)

func NewRouter(cfg *config.Config) (*echo.Echo, error) {
	router := echo.New()
	router.HideBanner = true

	router.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "Hello, Wave! 🌊")
		return nil
	})

	return router, nil
}
