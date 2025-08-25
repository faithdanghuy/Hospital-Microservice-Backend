package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/log"
	coreModel "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/golang-jwt/jwt"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// jwtKeyFunc returns the key for validating JWT tokens.
func jwtKeyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != "HS256" {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}
	return []byte(viper.GetString("JWT_SECRET_KEY")), nil
}

// parseJWT parses and validates the JWT token.
func parseJWT(auth string) (*jwt.Token, error) {
	token, err := jwt.Parse(auth, jwtKeyFunc)
	if err != nil {
		return nil, fmt.Errorf("parseJWT: error parse token: %w", err)
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

// setCustomClaims sets custom JWT claims to the echo context.
func setCustomClaims(c echo.Context, token *jwt.Token) {
	claims, _ := token.Claims.(jwt.MapClaims)
	var (
		id          string
		accountType string
	)
	if v, ok := claims["id"].(string); ok {
		id = v
	}
	if v, ok := claims["account_type"].(string); ok {
		accountType = v
	}

	c.Set("user", coreModel.JwtCustomClaims{
		ID:          id,
		AccountType: accountType,
	})
}

// JWT returns a middleware function for handling JWT authentication.
func JWT() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(viper.GetString("JWT_SECRET_KEY")),
		ParseTokenFunc: func(c echo.Context, tokenString string) (interface{}, error) {
			return parseJWT(tokenString)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			log.Error("JWT.ErrorHandler error parsing token", zap.Error(err))
			return response.Error(c, http.StatusUnauthorized)
		},
		SuccessHandler: func(c echo.Context) {
			token, _ := c.Get("user").(*jwt.Token)
			setCustomClaims(c, token)
		},
	}
	return echojwt.WithConfig(config)
}
