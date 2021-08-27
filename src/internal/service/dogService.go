package service

import (
	"Dogs/src/internal/models"
	"Dogs/src/internal/repository"
	"context"
	"errors"
	"log"
)

type DogService struct {
	dogRepo repository.DogPostgresRepository
}

func (p *DogService) CreateDog(ctx  context.Context,dog models.Dog) error{

	if dog.DogName == ""{
		return errors.New("Field name is empty!")
	}

	err := p.dogRepo.AddDog(ctx,dog);if err != nil{
		log.Println(err)
	}
	return err
}

func NewDogService(dogRepo repository.DogPostgresRepository) *DogService {
	return &DogService{dogRepo: dogRepo}
}