package service

import (
	"context"
	"log"
	"main/internal/models"
	"main/internal/repository"
)

type DogService struct {
	dogRepo repository.DogPostgresRepository
}

func (p *DogService) Create(ctx context.Context, dog models.Dog) error {

	err := p.dogRepo.Add(ctx, dog)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *DogService) Get(ctx context.Context, name string) (models.Dog, error) {

	resultDog, err := p.dogRepo.Get(ctx, name)
	if err != nil {
		log.Println(err)
	}

	return resultDog, err
}

func (p *DogService) Change(ctx context.Context, name string, dog models.Dog) error {

	err := p.dogRepo.Change(ctx, name, dog)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (p *DogService) Delete(ctx context.Context, name string) error {

	err := p.dogRepo.Delete(ctx, name)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func NewDogService(dogRepo repository.DogPostgresRepository) *DogService {
	return &DogService{dogRepo: dogRepo}
}
