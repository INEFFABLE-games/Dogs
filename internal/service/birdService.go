package service

import (
	"context"
	"errors"
	"main/internal/models"
	"main/internal/repository"
)

type BirdService struct {
	birdRepo repository.BirdMongoRepository
}

func (b *BirdService) Create(ctx context.Context, bird models.Bird, token string) error {

	resultToken, err := b.birdRepo.ValidateToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	err = b.birdRepo.Create(ctx, bird, resultToken.Login)
	if err != nil {
		return err
	}

	return err
}

func (b *BirdService) Get(ctx context.Context, token string) (models.Bird, error) {

	resultToken, err := b.birdRepo.ValidateToken(ctx, token)
	if err != nil {
		return models.Bird{}, errors.New("unable to validate token")
	}

	resultBird, err := b.birdRepo.Get(ctx, resultToken.Login)
	if err != nil {
		return resultBird, err
	}

	return resultBird, err
}

func (b *BirdService) Change(ctx context.Context, bird models.Bird, token string) error {

	resultToken, err := b.birdRepo.ValidateToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	err = b.birdRepo.Change(ctx, resultToken.Login, bird)

	return err
}

func (b *BirdService) Delete(ctx context.Context, token string) error {

	resultToken, err := b.birdRepo.ValidateToken(ctx, token)
	if err != nil {
		return errors.New("unable to validate token")
	}

	err = b.birdRepo.Delete(ctx, resultToken.Login)

	return err
}

func NewBirdService(birdRepo repository.BirdMongoRepository) *BirdService {
	return &BirdService{birdRepo: birdRepo}
}
