package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/controllers/domain"
	"github.com/Wave-ETH-Global/wave-node/middlewares"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/logger"
	"github.com/pkg/errors"
)

var ErrRepository = errors.New("repository")

const (
	TokenExpirationTime = time.Hour * 24
)

type LoginService struct {
	cfg *config.Config
	rc  *redis.Client
}

func NewLoginService(cfg *config.Config, rc *redis.Client) *LoginService {
	return &LoginService{
		cfg: cfg,
	}
}

func (ls *LoginService) VerifyWalletSignature(walletAddress string, signature string, nonce string) error {
	walletAddressDomain, err := domain.NewAddress(walletAddress)
	if err != nil {
		logger.Errorf("invalid wallet address:%s", walletAddress)
		return errors.Wrap(ErrRepository, err.Error())
	}
	loginMessage := signatureSourceMessageFromTemplate(domain.WalletSignatureMessage, nonce)
	if err := walletAddressDomain.SigVerify(loginMessage, signature); err != nil {
		logger.Errorf("failed to verify sig. sig:%s, loginMessage:%s", signature, loginMessage)
		return errors.WithStack(err)
	}

	return nil
}

func signatureSourceMessageFromTemplate(messageTemplate string, nonce string) string {
	return fmt.Sprintf(string(messageTemplate), nonce)
}

func (ls *LoginService) IssueJWT(address string, uuid string) (string, error) {
	claims := middlewares.JWTCustomClaims{
		uuid,
		address,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpirationTime).Unix(),
		},
	}

	// generate header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token
	tokenString, err := token.SignedString(ls.cfg.JWT.SigningKey)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256([]byte(tokenString))
	ls.rc.Set(context.TODO(), hex.EncodeToString(hash[:]), true, TokenExpirationTime)
	return hex.EncodeToString(hash[:]), nil
}
