package controllers

import (
	"context"
	"crypto/sha256"
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

func NewLoginController(rc *redis.Client, ec *ethclient.EthClient) *LoginController {
	return &LoginController{rc: rc, ec: ec}
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
	ctx := ectx.Request().Context()
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
		logger.Error(err)
		return err
	}

	if ac == nil {
		logger.Errorf("account with address %s isn't found", req.WalletAddress)
	}

	claims := jwt.MapClaims{
		"address": req.WalletAddress,
		"uuid":    ac.UUID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// generate header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// add sig to token
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return err
	}
	ectx.Response().Header().Set("Authorization", tokenString)
	hash := sha256.Sum256([]byte(tokenString))
	c.rc.Set(ctx, hex.EncodeToString(hash[:]), true, time.Duration(time.Hour*24*7))

	return ectx.JSON(http.StatusOK, &LoginByWalletAddressRes{Token: tokenString})
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

	newUUID := uuid.New().String()
	p.UUID = newUUID
	pj, _ := json.Marshal(p)

	r, _ := utils.SecureRandom(24)
	err = lc.rc.Set(context.TODO(), r, string(pj), time.Hour*3).Err()
	if err != nil {
		return err
	}

	var resp struct {
		UUID string `json:"uuid"`
	}

	resp.UUID = newUUID

	return ctx.JSON(http.StatusOK, resp)
}

func (lc *LoginController) SignupCompleted(ctx echo.Context) error {
	var t struct {
		OneTimeToken string `json:"one_time_token"`
	}

	if err := ctx.Bind(&t); err != nil {
		return err
	}
	req := lc.rc.Get(context.TODO(), t.OneTimeToken)
	if req.Err() != nil {
		return req.Err()
	}

	var p models.Profile
	if err := json.Unmarshal([]byte(req.String()), &p); err != nil {
		return err
	}

	uuid, err := lc.ec.GetUUIDOfClaimedUserHandle(p.Username)
	if err != nil {
		return err
	}
	if uuid != p.UUID {
		return errors.New("uuid doesn't match with what we have given to user")
	}

	if err := lc.repo.InsertProfile(&p); err != nil {
		return err
	}

	return nil
}
