package service

import (
	"context"
	"errors"
	"main/internal/models"
	"main/internal/repository"
)

// DogService create new dog service.
type DogService struct {
	dogRepo repository.DogRepository
}

// Create calls add func on dog repository.
func (p *DogService) Create(ctx context.Context, dog models.Dog, token string) error {

	result, err := p.dogRepo.ValidateToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	return p.dogRepo.Create(ctx, dog, result.Login)
}

// ValidateToken check is token exists.
func (p *DogService) ValidateToken(ctx context.Context, token string) (models.Token, error) {
	return p.dogRepo.ValidateToken(ctx, token)
}

// Get call get func on dog repository.
func (p *DogService) Get(ctx context.Context, token string) (models.Dog, error) {

	result, err := p.dogRepo.ValidateToken(ctx, token)
	if err != nil {
		return models.Dog{}, errors.New("unable to validate token")
	}

	return p.dogRepo.Get(ctx, result.Login)
}

// Change call change func on dog repository.
func (p *DogService) Change(ctx context.Context, token string, dog models.Dog) error {
	result, err := p.dogRepo.ValidateToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	return p.dogRepo.Change(ctx, result.Login, dog)
}

// Delete call delete func on dog repository.
func (p *DogService) Delete(ctx context.Context, token string) error {
	result, err := p.dogRepo.ValidateToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	// check is dog exist
	_, err = p.dogRepo.Get(ctx, result.Login)
	if err != nil {
		return errors.New("dog doesn't exist")
	}

	return p.dogRepo.Delete(ctx, result.Login)
}

// NewDogService creates new dog service.
func NewDogService(dogRepo repository.DogRepository) *DogService {
	return &DogService{dogRepo: dogRepo}
}
