package middlewares

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/utils"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	UserContext = "user_context"
)

type AuthMiddleware struct {
	rc  *redis.Client
	cfg *config.Config
}

func NewAuthMiddleware(rc *redis.Client, cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{rc: rc, cfg: cfg}
}

func (am *AuthMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
		claims, err := am.CheckToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.HTTPError{Error: err.Error()})
			return err
		}
		c.Set(UserContext, claims)
		return next(c)
	}
}

func (am *AuthMiddleware) CheckToken(token string) (*JWTCustomClaims, error) {
	if token == "" {
		return nil, errors.New("no auth token provided!")
	}
	hash := sha256.Sum256([]byte(token))
	ok, err := am.rc.Get(context.Background(), hex.EncodeToString(hash[:])).Bool()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New("token doesn't exist in session storage")
		}
		return nil, err
	}

	if !ok {
		return nil, errors.New("token is invalid")
	}

	var claims JWTCustomClaims
	tok, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(am.cfg.JWT.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !tok.Valid {
		return nil, errors.New("token is invalid!")
	}

	return tok.Claims.(*JWTCustomClaims), nil
}
