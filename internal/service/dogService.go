package service

import (
	"context"
	"main/internal/models"
	"main/internal/repository"
)

// DogService create new dog service.
type DogService struct {
	dogRepo repository.DogRepository
}

// Create calls add func on dog repository.
func (p *DogService) Create(ctx context.Context, dog models.Dog) error {
	err := p.dogRepo.Add(ctx, dog)

	return err
}

// Get call get func on dog repository.
func (p *DogService) Get(ctx context.Context, name string) (models.Dog, error) {
	resultDog, err := p.dogRepo.Get(ctx, name)

	return resultDog, err
}

// Change call change func on dog repository.
func (p *DogService) Change(ctx context.Context, name string, dog models.Dog) error {
	err := p.dogRepo.Change(ctx, name, dog)

	return err
}

// Delete call delete func on dog repository.
func (p *DogService) Delete(ctx context.Context, name string) error {
	err := p.dogRepo.Delete(ctx, name)

	return err
}

// NewDogService creates new dog service.
func NewDogService(dogRepo repository.DogRepository) *DogService {
	return &DogService{dogRepo: dogRepo}
}
