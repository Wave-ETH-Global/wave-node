package controllers

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/repositories"
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
		ctx.JSON(http.StatusInternalServerError, HTTPError{err.Error()})
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}
