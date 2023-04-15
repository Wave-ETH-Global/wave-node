package controllers

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/repositories"
	"github.com/google/logger"
	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	pr *repositories.ProfileRepository
}

func NewProfileController(pr *repositories.ProfileRepository) *ProfileController {
	return &ProfileController{pr: pr}
}

type RequestConnectionReq struct {
	RequestedAddress string   `json:"requested_address"`
	Tags             []string `json:"tags"`
	PrivateTags      []string `json:"private_tags"`
}

func (pc *ProfileController) RequestConnection(ctx echo.Context) error {
	req := &RequestConnectionReq{}
	c := ctx.Request().Context()
	connectorAddress := ctx.Param("address")
	if err := ctx.Bind(req); err != nil {
		return err
	}

	logger.Infof("req:%+v, wallet:%s", req, connectorAddress)

	err := pc.pr.RequestConnection(
		c,
		repositories.RequestConnectionParam{
			ConnectorAddress: connectorAddress,
			RequestedAddress: req.RequestedAddress,
			Tags:             req.Tags,
			PrivateTags:      req.PrivateTags,
		},
	)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (pc *ProfileController) GetProfileChainInfo(ctx echo.Context) error {
	data, err := pc.pr.GetOnchainInfo(ctx.Param("address"), ctx.QueryParam("cursor"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{err.Error()})
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}
