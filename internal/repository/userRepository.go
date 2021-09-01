package repository

import (
	"context"
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
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

// InsertToken create new token and execute query to add token in database.
func (u *UserRepository) InsertToken(ctx context.Context, login string, token string) (string, error) {

	_, err := u.db.ExecContext(ctx,
		"insert into tokens(login,value) values($1,$2)",
		login,
		token,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

// UpdateToken create new token and execute query to add token in database.
func (u *UserRepository) UpdateToken(ctx context.Context, login string, token string) (string, error) {

	_, err := u.db.ExecContext(ctx,
		"update tokens set value=$1 where login=$2",
		token,
		login,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

// CheckToken create new token and execute query to add token in database.
func (u *UserRepository) CheckToken(ctx context.Context, login string) (bool, error) {

	tok := models.Token{}

	err := u.db.QueryRowContext(ctx,
		"select * from tokens where login=$1",
		login,
	).Scan(&tok.Login, &tok.Value)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Login check is user exist and create new token for user.
func (u *UserRepository) Login(ctx context.Context, usr models.User, token string) (string, error) {
	_, err := u.Get(ctx, usr)
	if err != nil {
		return "", err
	}

	isTokenAlreadyExists, err := u.CheckToken(ctx, usr.Login)
	if !isTokenAlreadyExists {
		token, err = u.InsertToken(ctx, usr.Login, token)
		if err != nil {
			return "", err
		}
	}

	token, err = u.UpdateToken(ctx, usr.Login, token)

	return token, nil
}

// Logout check is token exists and execute query to delete token from database.
func (u *UserRepository) Logout(ctx context.Context, token string) error {
	res, err := u.db.ExecContext(ctx,
		"select * from tokens where value=$1",
		token,
	)
	if err != nil {
		return err
	}

	affectedRow, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRow == 0 {
		log.WithFields(log.Fields{
			"handler": "user",
			"action":  "logout",
		}).Debug("doesn't find token")

		return errors.New("doesn't find token")
	}

	_, err = u.db.ExecContext(ctx,
		"delete from tokens where value=$1",
		token,
	)

	return err
}

// NewUserRepository create new user repository object.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
