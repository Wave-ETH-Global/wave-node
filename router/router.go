package router

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(cfg *config.Config,
	pc *controllers.ProfileController,
	ac *controllers.AccountController,
) (*echo.Echo, error) {
	router := echo.New()
	router.HideBanner = true

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	router.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "Hello, Wave! 🌊")
		return nil
	})

	router.POST("/token", ac.LoginByOrRegisterWalletAddress)
	router.GET("/profile/chaininfo/:address", pc.GetProfileChainInfo)

	router.POST("/profile/:address/connections", pc.RequestConnection)

	return router, nil
}
