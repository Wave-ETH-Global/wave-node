package controllers

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/middlewares"
	"github.com/Wave-ETH-Global/wave-node/models"
	"github.com/Wave-ETH-Global/wave-node/repositories"
	"github.com/Wave-ETH-Global/wave-node/utils"
	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	pr *repositories.ProfileRepository
}

func NewProfileController(pr *repositories.ProfileRepository) *ProfileController {
	return &ProfileController{pr: pr}
}

func (pc *ProfileController) GetProfileChainInfo(ctx echo.Context) error {
	data, err := pc.pr.GetOnchainInfo(ctx.Param("address"), ctx.QueryParam("cursor"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HTTPError{Error: err.Error()})
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (pc *ProfileController) GetProfileConnections(ctx echo.Context) error {
	data, err := pc.pr.GetProfileConnections(ctx.Get(middlewares.UserContext).(*middlewares.JWTCustomClaims).UUID)
	if err != nil {
		return err
	}

	var resp struct {
		Connections []*models.Connection `json:"connections"`
	}

	resp.Connections = data
	return ctx.JSON(http.StatusOK, resp)
}
