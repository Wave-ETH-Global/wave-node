package middlewares

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	UUID       string `json:"uuid"`
	ETHAddress string `json:"eth_address"`
	jwt.StandardClaims
}
