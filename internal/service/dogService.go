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
	return p.dogRepo.Add(ctx, dog)
}

// Get call get func on dog repository.
func (p *DogService) Get(ctx context.Context, name string) (models.Dog, error) {
	return p.dogRepo.Get(ctx, name)
}

// Change call change func on dog repository.
func (p *DogService) Change(ctx context.Context, name string, dog models.Dog) error {
	return p.dogRepo.Change(ctx, name, dog)
}

// Delete call delete func on dog repository.
func (p *DogService) Delete(ctx context.Context, name string) error {
	return p.dogRepo.Delete(ctx, name)
}

// NewDogService creates new dog service.
func NewDogService(dogRepo repository.DogRepository) *DogService {
	return &DogService{dogRepo: dogRepo}
}
