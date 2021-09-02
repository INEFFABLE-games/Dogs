package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"main/internal/models"
	"main/internal/repository"
	"time"
)

type AuthService struct {
	tokenRepo repository.TokenRepository
}

// RefreshTokens refresh existing tokens for user
func (a *AuthService) RefreshTokens(ctx context.Context, userLogin string) (string, string, error) {

	//JWT claims
	claims := models.CustomClaims{
		Login: userLogin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "",
		},
	}

	jwtTok, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("dog"))
	if err != nil {
		return "", "", err
	}

	//RT claims
	claims = models.CustomClaims{
		Login: userLogin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "",
		},
	}

	rt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("RefTokKey"))

	err = a.tokenRepo.InsertOrUpdate(ctx, rt, userLogin)

	return jwtTok, rt, err
}

// NewAuthService creates new AuthService object
func NewAuthService(tokenRepo repository.TokenRepository) *AuthService {
	return &AuthService{tokenRepo: tokenRepo}
}
