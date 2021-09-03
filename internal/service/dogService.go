package service

import (
	"Dogs/internal/models"
	"Dogs/internal/repository"
	"context"
	"errors"
)

// DogService create new dog service.
type DogService struct {
	dogRepo repository.DogRepository
}

// Create calls add func on dog repository.
func (p *DogService) Create(ctx context.Context, dog models.Dog) error {
	return p.dogRepo.Create(ctx, dog)
}

// CheckToken check is token exists.
func (p *DogService) CheckToken(ctx context.Context, token string) (models.Token, error) {
	return p.dogRepo.CheckToken(ctx, token)
}

// Get call get func on dog repository.
func (p *DogService) Get(ctx context.Context, userLogin string) (models.Dog, error) {
	return p.dogRepo.Get(ctx, userLogin)
}

// Change call change func on dog repository.
func (p *DogService) Change(ctx context.Context, dog models.Dog) error {
	return p.dogRepo.Change(ctx, dog)
}

// Delete call delete func on dog repository.
func (p *DogService) Delete(ctx context.Context, userLogin string) error {
	// check is dog exist
	_, err := p.dogRepo.Get(ctx, userLogin)
	if err != nil {
		return errors.New("dog doesn't exist")
	}

	return p.dogRepo.Delete(ctx, userLogin)
}

// NewDogService creates new dog service.
func NewDogService(dogRepo repository.DogRepository) *DogService {
	return &DogService{dogRepo: dogRepo}
}
