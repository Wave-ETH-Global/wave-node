package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/logger"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/Wave-ETH-Global/wave-node/controllers/domain"
	"github.com/Wave-ETH-Global/wave-node/middlewares"
	"github.com/Wave-ETH-Global/wave-node/repositories"
)

type AccountController struct {
	repo    *repositories.AccountRepo
	session *middlewares.SessionStore
}

func NewAccountController(repo *repositories.AccountRepo, session *middlewares.SessionStore) *AccountController {
	return &AccountController{repo: repo, session: session}
}

type LoginByWalletAddressReq struct {
	WalletAddress string `json:"wallet_address"`
	Signature     string `json:"signature"`
	Nonce         string `json:"nonce"`
}

type LoginByWalletAddressRes struct {
	WalletAddress string `json:"wallet_address"`
}

func (c AccountController) LoginByOrRegisterWalletAddress(
	ectx echo.Context,
) error {
	req := &LoginByWalletAddressReq{}
	ctx := ectx.Request().Context()
	if err := ectx.Bind(req); err != nil {
		return err
	}

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

	ac := c.repo.GetAccountByAddress(ctx, domain.Address(req.WalletAddress))

	if ac == nil {
		c.repo.CreateAddress(ctx, walletAddress)
	}
	if ac != nil {
		logger.Infof("found an account. wallet_address:%s", req.WalletAddress)
	}

	claims := jwt.MapClaims{
		"address": req.WalletAddress,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// generate header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// add sig to token
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return err
	}
	ectx.Response().Header().Set("authorization", tokenString)
	c.session.SetWalletAddress(ctx, walletAddress)

	return ectx.JSON(http.StatusOK, &LoginByWalletAddressRes{WalletAddress: req.WalletAddress})
}

func SignatureSourceMessageFromTemplate(messageTemplate string, nonce string) string {
	return fmt.Sprintf(string(messageTemplate), nonce)
}
