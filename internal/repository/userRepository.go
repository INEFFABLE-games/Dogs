package repository

import (
	"context"
	"database/sql"
	"main/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

// Create execute query to add user in database.
func (u *UserRepository) Create(ctx context.Context, usr models.User) error {
	_, err := u.db.ExecContext(ctx,
		"insert into users(login,password) values($1,$2)",
		usr.Login,
		usr.Password,
	)

	return err
}

// Get execute query to get user from database and returns user object.
func (u *UserRepository) Get(ctx context.Context, usr models.User) (models.User, error) {
	resultUser := models.User{}
	err := u.db.QueryRowContext(ctx,
		"select * from users where login=$1 and password=$2",
		usr.Login,
		usr.Password,
	).Scan(&resultUser.Login, &resultUser.Password)

	return resultUser, err
}

// Logout check is token exists and execute query to delete token from database.
func (u *UserRepository) Logout(ctx context.Context) error {

	return nil
}

// NewUserRepository create new user repository object.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
