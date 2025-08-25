package security

import (
	"fmt"
	"time"

	coreModel "github.com/Hospital-Microservice/hospital-core/model"
	"github.com/golang-jwt/jwt/v5"

	"github.com/spf13/viper"
)

type Token struct {
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
}

func GenToken(ID string, AccountType string, duration time.Duration) (*string, error) {
	var (
		claimsToken = &coreModel.JwtCustomClaims{
			ID:          ID,
			AccountType: AccountType,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			},
		}
		secretKey      = []byte(viper.Get("JWT_SECRET_KEY").(string))
		accessTokenObj = jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	)

	var accessTokenString, err = accessTokenObj.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &accessTokenString, nil
}

func ParseTokenFromString(tokenString string) (*coreModel.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&coreModel.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.Get("JWT_SECRET_KEY").(string)), nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*coreModel.JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
