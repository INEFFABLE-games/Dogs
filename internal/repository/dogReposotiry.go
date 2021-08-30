package repository

import (
	"context"
	"database/sql"
	"main/internal/models"
)

// DogRepository creates new dogs repository.
type DogRepository struct {
	db *sql.DB
}

// Add insert add query in db.
func (d *DogRepository) Add(ctx context.Context, dog models.Dog) error {
	_, err := d.db.ExecContext(ctx, "insert into dogs(Name,Gender) values($1,$2)", dog.Name, dog.Gender)

	return err
}

// Get insert select query in db and return dog object.
func (d *DogRepository) Get(ctx context.Context, name string) (models.Dog, error) {
	resultDog := models.NewDog("", "")
	err := d.db.QueryRowContext(ctx, "select * from dogs where name = $1", name).Scan(&resultDog.Name, &resultDog.Gender)

	return resultDog, err
}

// Change insert change query in db.
func (d *DogRepository) Change(ctx context.Context, name string, dog models.Dog) error {
	_, err := d.db.ExecContext(ctx, "update dogs set name = $1, gender = $2 where name = $3", dog.Name, dog.Gender, name)

	return err
}

// Delete insert delete query in db.
func (d *DogRepository) Delete(ctx context.Context, name string) error {
	_, err := d.db.ExecContext(ctx, "delete from dogs where name = $1", name)

	return err
}

// NewDogRepository creates new repository for dogs.
func NewDogRepository(db *sql.DB) DogRepository {
	return DogRepository{db: db}
}
