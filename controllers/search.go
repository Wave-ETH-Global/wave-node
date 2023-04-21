package controllers

import (
	"net/http"

	"github.com/Wave-ETH-Global/wave-node/models"
	"github.com/Wave-ETH-Global/wave-node/repositories"
	"github.com/labstack/echo/v4"
)

type SearchController struct {
	pr *repositories.ProfileRepository
}

func NewSearchController(pr *repositories.ProfileRepository) *SearchController {
	return &SearchController{pr: pr}
}

func (sc *SearchController) Search(ctx echo.Context) error {
	var req struct {
		Tags []string `json:"tags"`
	}
	ctx.Bind(&req)

	data, err := sc.pr.SearchProfiles(ctx.QueryParam("uuid"), req.Tags)

	var resp struct {
		Acquantancies []*models.Profile `json:"acquantancies"`
	}

	if err != nil {
		return err
	}

	resp.Acquantancies = data

	return ctx.JSON(http.StatusOK, resp)
}
