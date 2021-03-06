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
	return u.userRepo.Login(ctx, usr)
}

// Logout calls logout func on user repository.
func (u *UserService) Logout(ctx context.Context, token string) error {
	return u.userRepo.Logout(ctx, token)
}

// NewUserService creates new user service object.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
