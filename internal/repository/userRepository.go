package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"time"

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

// CreateToken create new token and execute query to add token in database.
func (u *UserRepository) CreateToken(ctx context.Context, login string) (string, error) {
	claims := models.CustomClaims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
			Issuer:    "",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SigningString()
	if err != nil {
		return "", err
	}

	_, err = u.db.ExecContext(ctx,
		"insert into tokens(login,value) values($1,$2)",
		login,
		token,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

// Login check is user exist and create new token for user.
func (u *UserRepository) Login(ctx context.Context, usr models.User) (string, error) {
	_, err := u.Get(ctx, usr)
	if err != nil {
		return "", err
	}

	token, err := u.CreateToken(ctx, usr.Login)
	if err != nil {
		return "", err
	}

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
