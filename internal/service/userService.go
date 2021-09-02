package service

import (
	"context"
	"main/internal/models"
	"main/internal/repository"
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

	_, err := u.userRepo.Get(ctx, usr)
	if err != nil {
		return "", err
	}

	token, err := CreateToken(usr.Login)
	if err != nil {
		return "", err
	}

	return token, err
}

// Logout calls logout func on user repository.
func (u *UserService) Logout(ctx context.Context) error {
	return u.userRepo.Logout(ctx)
}

// NewUserService creates new user service object.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
