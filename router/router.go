package router

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/controllers"
	"github.com/labstack/echo/v4"
)

func NewRouter(cfg *config.Config, pc *controllers.ProfileController) (*echo.Echo, error) {
	router := echo.New()
	router.HideBanner = true

	router.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "Hello, Wave! 🌊")
		return nil
	})

	router.GET("/profile/chaininfo/:address", pc.GetProfileChainInfo)

	return router, nil
}
