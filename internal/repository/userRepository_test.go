package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"main/internal/models"
)

const ctxtime = 5

func TestUserRepository_Create(t *testing.T) {
	base, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"hadler": "Test_UserRepository",
			"action": "Login",
		}).Errorf("unable connect to database %v", err)
	}

	r := NewUserRepository(base)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*ctxtime)
	defer cancel()

	err = r.Create(ctx, models.User{
		Login:    "testlogin",
		Password: "testpassword",
	})

	require.Nil(t, err)
}

func TestUserRepository_Login(t *testing.T) {
	base, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"hadler": "Test_UserRepository",
			"action": "Login",
		}).Errorf("unable connect to database %v", err)
	}

	r := NewUserRepository(base)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*ctxtime)
	defer cancel()

	_, err = r.GetByLogin(ctx, "testlogin")

	require.Nil(t, err)
}

func TestUserRepository_Logout(t *testing.T) {
	base, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"hadler": "Test_UserRepository",
			"action": "Login",
		}).Errorf("unable connect to database %v", err)
	}

	r := NewUserRepository(base)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*ctxtime)
	defer cancel()

	err = r.Logout(ctx)

	require.Nil(t, err)
}
