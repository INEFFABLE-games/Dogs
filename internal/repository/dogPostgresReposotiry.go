package repository

import (
	"Dogs/internal/models"
	"context"
	"database/sql"
	"log"
)

type DogPostgresRepository struct {
	db *sql.DB
}

func (d *DogPostgresRepository) Add(ctx context.Context, dog models.Dog) error {

	_, err := d.db.ExecContext(ctx, "insert into dogs(Name,Gender) values($1,$2)", dog.Name, dog.Gender)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (d *DogPostgresRepository) Get(ctx context.Context, dog models.Dog) (models.Dog, error) {

	resultDog := models.Dog{}
	err := d.db.QueryRowContext(ctx, "select * from dogs where name = $1", dog.Name).Scan(&resultDog.Name, &resultDog.Gender)
	if err != nil {
		log.Println(err)
	}

	return resultDog, err
}

func (d *DogPostgresRepository) Change(ctx context.Context, name string, dog models.Dog) error {

	_, err := d.db.ExecContext(ctx, "update dogs set name = $1, gender = $2 where name = $3", dog.Name, dog.Gender, name)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (d *DogPostgresRepository) Delete(ctx context.Context, name string) error {

	_, err := d.db.ExecContext(ctx, "delete from dogs where name = $1", name)
	if err != nil {
		log.Println(err)
	}

	return err
}

func NewDogPostgresRepository(db *sql.DB) DogPostgresRepository {
	return DogPostgresRepository{db: db}
}
