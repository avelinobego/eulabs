package jwt

import (
	"eulabs/products/internal/user/dtos"
	"time"

	"github.com/golang-jwt/jwt"
)

func Generate(secretKey string, login string) (reponse dtos.UserLoginResponse, err error) {

	expiresAt := time.Now().Add(10 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Subject:   login,
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	reponse = dtos.UserLoginResponse{
		AccessToken: tokenString,
		ExpiredAt:   expiresAt,
	}

	return
}
