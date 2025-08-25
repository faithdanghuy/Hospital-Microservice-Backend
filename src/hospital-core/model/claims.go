package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID          string `json:"id" mapstructure:"id"`
	AccountType string `json:"account_type" mapstructure:"account_type"`
	jwt.RegisteredClaims
}
