package repository

import (
	"context"
	"database/sql"
)

// TokenRepository struct for tokens repository
type TokenRepository struct {
	db *sql.DB
}

/*
func (t *TokenRepository) Delete(ctx context.Context,userLogin string) error{
	_,err := t.db.ExecContext(ctx,
		"delete from tokens where login=$1",
		userLogin,
	)

	return err
}*/

// InsertRT insert new token in tokens table
func (t *TokenRepository) InsertOrUpdate(ctx context.Context, RT string, userLogin string) error {

	_, err := t.db.ExecContext(ctx,
		"insert into tokens(login,value) values($1,$2) ON CONFLICT (login) DO update set login=$1, value=$2",
		userLogin,
		RT,
	)

	return err
}

// NewTokenRepository creates new TokenRepository object
func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{db: db}
}
