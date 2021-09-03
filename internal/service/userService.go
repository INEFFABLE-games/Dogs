package service

import (
	"Dogs/internal/models"
	"Dogs/internal/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

// UserService create new user service object.
type UserService struct {
	userRepo repository.UserRepository
}

// Create calls create func on user repository.
func (u *UserService) Create(ctx context.Context, usr models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usr.Password = string(hashedPassword)

	return u.userRepo.Create(ctx, usr)
}

// Login call login func on user repository.
func (u *UserService) Login(ctx context.Context, usr models.User) error {

	userFromDatabase, err := u.userRepo.GetByLogin(ctx, usr.Login)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDatabase.Password), []byte(usr.Password))
	if err != nil {
		return err
	}

	return err
}

// Logout calls logout func on user repository.
func (u *UserService) Logout(ctx context.Context) error {
	return u.userRepo.Logout(ctx)
}

// NewUserService creates new user service object.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
