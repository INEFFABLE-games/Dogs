package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"main/internal/models"
	"main/internal/repository"
	"time"
)

// UserService create new user service object.
type UserService struct {
	userRepo repository.UserRepository
}

// Create calls create func on user repository.
func (u *UserService) Create(ctx context.Context, usr models.User) error {
	return u.userRepo.Create(ctx, usr)
}

// Login call login func on user repository.
func (u *UserService) Login(ctx context.Context, usr models.User) (string, error) {

	claims := models.CustomClaims{
		Login: usr.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("dog"))
	if err != nil {
		return "", err
	}

	return u.userRepo.Login(ctx, usr, token)
}

// Logout calls logout func on user repository.
func (u *UserService) Logout(ctx context.Context, token string) error {
	return u.userRepo.Logout(ctx, token)
}

// NewUserService creates new user service object.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
