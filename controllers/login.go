package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/logger"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/Wave-ETH-Global/wave-node/controllers/domain"
	"github.com/Wave-ETH-Global/wave-node/middlewares"
	"github.com/Wave-ETH-Global/wave-node/repositories"
)

type LoginController struct {
	rc   *redis.Client
	repo *repositories.ProfileRepository
}

func NewLoginController(rc *redis.Client) *LoginController {
	return &LoginController{rc: rc}
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
