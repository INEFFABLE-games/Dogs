package repository

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/models"
)

// DogRepository creates new dog repository.
type DogRepository struct {
	db *sql.DB
}

// ValidateToken check is token exists.
func (d *DogRepository) ValidateToken(ctx context.Context, token string) (models.Token, error) {

	result := models.Token{}

	err := d.db.QueryRowContext(ctx,
		"select * from tokens where value=$1",
		token,
	).Scan(&result.Login, &result.Value)

	if err != nil {
		return result, errors.New("unable to get token")
	}

	return result, err
}

// Create insert add query in db.
func (d *DogRepository) Create(ctx context.Context, dog models.Dog, owner string) error {
	_, err := d.db.ExecContext(ctx,
		"insert into dogs(Owner,Name,Gender) values($1,$2,$3)",
		owner,
		dog.Name,
		dog.Gender,
	)

	return err
}

// Get insert select query in db and return dog object.
func (d *DogRepository) Get(ctx context.Context, owner string) (models.Dog, error) {
	resultDog := models.Dog{}
	err := d.db.QueryRowContext(ctx,
		"select * from dogs where owner = $1",
		owner,
	).Scan(&resultDog.Owner, &resultDog.Name, &resultDog.Gender)

	return resultDog, err
}

// Change insert change query in db.
func (d *DogRepository) Change(ctx context.Context, owner string, dog models.Dog) error {
	_, err := d.db.ExecContext(ctx,
		"update dogs set name = $1, gender = $2 where owner = $3",
		dog.Name,
		dog.Gender,
		owner,
	)

	return err
}

// Delete insert delete query in db.
func (d *DogRepository) Delete(ctx context.Context, owner string) error {
	_, err := d.db.ExecContext(ctx,
		"delete from dogs where owner = $1",
		owner,
	)

	return err
}

// NewDogRepository creates new repository for dogs.
func NewDogRepository(db *sql.DB) *DogRepository {
	return &DogRepository{db: db}
}
