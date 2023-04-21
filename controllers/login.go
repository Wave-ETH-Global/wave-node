package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/Wave-ETH-Global/wave-node/ethclient"
	"github.com/Wave-ETH-Global/wave-node/models"
	"github.com/Wave-ETH-Global/wave-node/repositories"
	"github.com/Wave-ETH-Global/wave-node/services"
	"github.com/Wave-ETH-Global/wave-node/utils"
)

type LoginController struct {
	rc   *redis.Client
	repo *repositories.ProfileRepository
	ec   *ethclient.EthClient
	ls   *services.LoginService
}

func NewLoginController(rc *redis.Client, ec *ethclient.EthClient, pr *repositories.ProfileRepository) *LoginController {
	return &LoginController{rc: rc, ec: ec, repo: pr}
}

type LoginByWalletAddressReq struct {
	WalletAddress string `json:"wallet_address"`
	Signature     string `json:"signature"`
	Nonce         string `json:"nonce"`
}

type LoginByWalletAddressRes struct {
	Token string `json:"token"`
}

func (c *LoginController) LoginByWallet(
	ectx echo.Context,
) error {
	req := &LoginByWalletAddressReq{}
	if err := ectx.Bind(req); err != nil {
		return err
	}

	logger.Infof("req:%+v", req)

	err := c.ls.VerifyWalletSignature(req.WalletAddress, req.Signature, req.Nonce)
	if err != nil {
		return err
	}

	ac, err := c.repo.GetProfileByAddress(req.WalletAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return ectx.JSON(http.StatusBadRequest, utils.HTTPError{Error: "no such profile"})
		}
		logger.Error(err)
		return err
	}

	token, err := c.ls.IssueJWT(req.WalletAddress, ac.UUID)
	if err != nil {
		logger.Error(err)
		return err
	}

	ectx.Response().Header().Set("Authorization", token)

	return ectx.JSON(http.StatusOK, &LoginByWalletAddressRes{Token: token})
}

func (lc *LoginController) Signup(ctx echo.Context) error {
	p := models.Profile{}
	err := ctx.Bind(&p)
	if err != nil {
		logger.Error(err)
		return err
	}

	if p.ETHAddress == "" {
		return ctx.JSON(http.StatusBadRequest, utils.HTTPError{Error: "you must specify ethereum address!"})
	}

	_, err = lc.repo.GetProfileByAddress(p.ETHAddress)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, utils.HTTPError{Error: "profile with such ethereum address exists!"})
	} else if err != sql.ErrNoRows && err != nil {
		logger.Error(err)
		return err
	}

	newUUID := uuid.New().String()
	p.UUID = newUUID
	pj, _ := json.Marshal(p)

	r, _ := utils.SecureRandom(24)
	err = lc.rc.Set(context.TODO(), r, string(pj), time.Hour*3).Err()
	if err != nil {
		logger.Error(err)
		return err
	}

	var resp struct {
		UUID         string `json:"uuid"`
		OneTimeToken string `json:"one_time_token"`
	}

	resp.UUID = newUUID
	resp.OneTimeToken = r

	return ctx.JSON(http.StatusOK, resp)
}

func (lc *LoginController) SignupCompleted(ctx echo.Context) error {
	var t struct {
		OneTimeToken string `json:"one_time_token"`
	}

	if err := ctx.Bind(&t); err != nil {
		logger.Error(err)
		return err
	}
	req := lc.rc.Get(context.TODO(), t.OneTimeToken)
	if req.Err() != nil {
		if req.Err() == redis.Nil {
			return ctx.JSON(http.StatusBadRequest, utils.HTTPError{Error: "one time token is invalid!"})
		}
		logger.Error(req.Err())
		return req.Err()
	}

	var p models.Profile
	if err := json.Unmarshal([]byte(req.Val()), &p); err != nil {
		logger.Error(err)
		return err
	}

	uuid, err := lc.ec.GetUUIDOfClaimedUserHandle(p.Username)
	if err != nil {
		logger.Error(err)
		return err
	}
	if uuid != p.UUID {
		return errors.New("uuid doesn't match with what we have given to user")
	}

	if err := lc.repo.InsertProfile(&p); err != nil {
		logger.Error(err)
		return err
	}

	lc.rc.Del(context.TODO(), t.OneTimeToken)

	token, err := lc.ls.IssueJWT(p.ETHAddress, p.UUID)
	if err != nil {
		logger.Error(err)
		return err
	}

	var resp struct {
		Token string `json:"token"`
	}

	resp.Token = token
	return ctx.JSON(http.StatusOK, resp)
}
