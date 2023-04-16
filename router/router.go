package router

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/controllers"
	"github.com/Wave-ETH-Global/wave-node/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(cfg *config.Config,
	pc *controllers.ProfileController,
	lc *controllers.LoginController,
	am *middlewares.AuthMiddleware,
	sc *controllers.SearchController,
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

	router.POST("/token", lc.LoginByWallet)
	router.GET("/profile/chaininfo/:address", pc.GetProfileChainInfo)
	router.POST("/signup", lc.Signup)
	router.POST("/signup/completed", lc.SignupCompleted)
	router.POST("/search", sc.Search)

	return router, nil
}
