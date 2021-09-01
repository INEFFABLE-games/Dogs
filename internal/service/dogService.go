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
func (p *DogService) Create(ctx context.Context, dog models.Dog, token string, userLogin string) error {

	_, err := p.CheckToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	return p.dogRepo.Create(ctx, dog, userLogin)
}

// CheckToken check is token exists.
func (p *DogService) CheckToken(ctx context.Context, token string) (models.Token, error) {
	return p.dogRepo.CheckToken(ctx, token)
}

// Get call get func on dog repository.
func (p *DogService) Get(ctx context.Context, token string, userLogin string) (models.Dog, error) {

	_, err := p.CheckToken(ctx, token)
	if err != nil {
		return models.Dog{}, errors.New("unable to validate token")
	}

	return p.dogRepo.Get(ctx, userLogin)
}

// Change call change func on dog repository.
func (p *DogService) Change(ctx context.Context, token string, dog models.Dog) error {
	resultToken, err := p.CheckToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	return p.dogRepo.Change(ctx, resultToken.Login, dog)
}

// Delete call delete func on dog repository.
func (p *DogService) Delete(ctx context.Context, token string, userLogin string) error {
	_, err := p.CheckToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	// check is dog exist
	_, err = p.dogRepo.Get(ctx, userLogin)
	if err != nil {
		return errors.New("dog doesn't exist")
	}

	return p.dogRepo.Delete(ctx, userLogin)
}

// NewDogService creates new dog service.
func NewDogService(dogRepo repository.DogRepository) *DogService {
	return &DogService{dogRepo: dogRepo}
}
