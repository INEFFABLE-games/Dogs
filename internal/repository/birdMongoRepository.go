package repository

import (
	"context"
	"database/sql"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/internal/models"
)

type BirdMongoRepository struct {
	mongodb   *mongo.Client
	postresdb *sql.DB
}

// ValidateToken check is token exists.
func (b *BirdMongoRepository) ValidateToken(ctx context.Context, token string) (models.Token, error) {

	result := models.Token{}

	err := b.postresdb.QueryRowContext(ctx,
		"select * from tokens where value=$1",
		token,
	).Scan(&result.Login, &result.Value)

	if err != nil {
		return result, errors.New("unable to get token")
	}

	return result, err
}

func (b *BirdMongoRepository) Create(ctx context.Context, bird models.Bird, owner string) error {

	bird.Owner = owner
	_, err := b.mongodb.Database("Birds").Collection("Birds").InsertOne(ctx, bird)
	if err != nil {
		return err
	}

	return err
}

func (b *BirdMongoRepository) Get(ctx context.Context, owner string) (models.Bird, error) {

	bird := models.Bird{}
	bird.Owner = owner

	err := b.mongodb.Database("Birds").Collection("Birds").FindOne(ctx, bson.M{"owner": owner}).Decode(&bird)
	if err != nil {
		return bird, err
	}

	return bird, err
}

func (b *BirdMongoRepository) Change(ctx context.Context, owner string, bird models.Bird) error {
	bird.Owner = owner
	updateData := bson.M{"$set": bird}
	_, err := b.mongodb.Database("Birds").Collection("Birds").UpdateOne(ctx, bson.M{"owner": owner}, updateData)
	if err != nil {
		return err
	}

	return err
}

func (b *BirdMongoRepository) Delete(ctx context.Context, owner string) error {

	_, err := b.mongodb.Database("Birds").Collection("Birds").DeleteOne(ctx, bson.M{"owner": owner})
	if err != nil {
		return err
	}

	return err
}

func NewBirdsMongoRepository(mongodb *mongo.Client, postgresdb *sql.DB) *BirdMongoRepository {
	return &BirdMongoRepository{
		mongodb:   mongodb,
		postresdb: postgresdb,
	}
}
