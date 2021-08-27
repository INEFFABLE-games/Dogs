package repository

import (
	"Dogs/src/internal/models"
	"context"
	"database/sql"
	"log"
)

type DogPostgresRepository struct {
	db *sql.DB
}

func (d *DogPostgresRepository) AddDog(ctx context.Context,dog models.Dog) error {

	_,err := d.db.ExecContext(ctx,"insert into dogs(DogName,Gender) values($1,$2)",dog.DogName,dog.Gender);if err != nil{
		log.Println(err)
	}
	return err
}

func NewDogPostgresRepository(db *sql.DB) DogPostgresRepository {
	return DogPostgresRepository{db: db}
}
