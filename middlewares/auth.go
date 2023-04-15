package middlewares

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	rc *redis.Client
}

func NewAuthMiddleware(rc *redis.Client) *AuthMiddleware {
	return &AuthMiddleware{rc: rc}
}

func (am *AuthMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := am.CheckToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			return err
		}
		return next(c)
	}
}

func (am *AuthMiddleware) CheckToken(token string) error {
	hash := sha256.Sum256([]byte(token))
	ok, err := am.rc.Get(context.Background(), hex.EncodeToString(hash[:])).Bool()
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("token is invalid")
	}

	return nil
}
