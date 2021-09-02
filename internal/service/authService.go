package service

import (
	"github.com/dgrijalva/jwt-go"
	"main/internal/models"
	"time"
)

func CreateToken(userLogin string) (string, error) {

	claims := models.CustomClaims{
		Login: userLogin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("dog"))
	if err != nil {
		return "", err
	}

	return token, err
}
