package controllers

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/Wave-ETH-Global/wave-node/controllers/domain"
	"github.com/Wave-ETH-Global/wave-node/ethclient"
	"github.com/Wave-ETH-Global/wave-node/middlewares"
	"github.com/Wave-ETH-Global/wave-node/models"
	"github.com/Wave-ETH-Global/wave-node/repositories"
	"github.com/Wave-ETH-Global/wave-node/utils"
)

type LoginController struct {
	rc   *redis.Client
	repo *repositories.ProfileRepository
	ec   *ethclient.EthClient
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

func (c LoginController) LoginByWallet(
	ectx echo.Context,
) error {
	req := &LoginByWalletAddressReq{}
	if err := ectx.Bind(req); err != nil {
		return err
	}

	logger.Infof("req:%+v", req)

	walletAddress, err := domain.NewAddress(req.WalletAddress)
	if err != nil {
		logger.Errorf("invalid wallet address:%s", req.WalletAddress)
		return errors.Wrap(middlewares.ErrRepository, err.Error())
	}

	sig := req.Signature
	loginMessage := SignatureSourceMessageFromTemplate(domain.WalletSignatureMessage, req.Nonce)
	if err := walletAddress.SigVerify(loginMessage, sig); err != nil {
		logger.Errorf("failed to verify sig. sig:%s, loginMessage:%s", sig, loginMessage)
		return errors.WithStack(err)
	}

	ac, err := c.repo.GetProfileByAddress(req.WalletAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return ectx.JSON(http.StatusBadRequest, HTTPError{Error: "no such profile"})
		}
		logger.Error(err)
		return err
	}

	token, err := c.issueJWT(req.WalletAddress, ac.UUID)
	if err != nil {
		logger.Error(err)
		return err
	}

	ectx.Response().Header().Set("Authorization", token)

	return ectx.JSON(http.StatusOK, &LoginByWalletAddressRes{Token: token})
}

func SignatureSourceMessageFromTemplate(messageTemplate string, nonce string) string {
	return fmt.Sprintf(string(messageTemplate), nonce)
}

func (lc *LoginController) Signup(ctx echo.Context) error {
	p := models.Profile{}
	err := ctx.Bind(&p)
	if err != nil {
		logger.Error(err)
		return err
	}

	if p.ETHAddress == "" {
		return ctx.JSON(http.StatusBadRequest, HTTPError{Error: "you must specify ethereum address!"})
	}

	_, err = lc.repo.GetProfileByAddress(p.ETHAddress)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, HTTPError{Error: "profile with such ethereum address exists!"})
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
			return ctx.JSON(http.StatusBadRequest, HTTPError{Error: "one time token is invalid!"})
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

	token, err := lc.issueJWT(p.ETHAddress, p.UUID)
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

func (lc *LoginController) issueJWT(address string, uuid string) (string, error) {
	claims := jwt.MapClaims{
		"address": address,
		"uuid":    uuid,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// generate header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// add sig to token
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256([]byte(tokenString))
	lc.rc.Set(context.TODO(), hex.EncodeToString(hash[:]), true, time.Duration(time.Hour*24*7))
	return hex.EncodeToString(hash[:]), nil
}
