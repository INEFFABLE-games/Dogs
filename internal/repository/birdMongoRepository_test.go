package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/internal/models"
	"testing"
	"time"
)

func initializeDataBases() (*mongo.Client, *sql.DB) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://dbuser1:dbuser1@petscluster.82emx.mongodb.net/Birds?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo
	mongoConn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "create",
		}).Errorf("unable connect to MonogDB %v", err)
	}

	//connect to postgres
	postgresconn, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"hadler": "Test_UserRepository",
			"action": "Login",
		}).Errorf("unable connect to database %v", err)
	}
	return mongoConn, postgresconn
}

func TestBirdsMongoRepository_Create(t *testing.T) {

	mongoConn, postgresConn := initializeDataBases()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	r := NewBirdsMongoRepository(mongoConn, postgresConn)

	err := r.Create(ctx, models.Bird{
		Owner: "",
		Name:  "Phoenix",
		Type:  "Firebird",
	}, "userOwner")

	require.Nil(t, err)
}

func TestBirdsMongoRepository_Get(t *testing.T) {
	mongoConn, postgresConn := initializeDataBases()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	r := NewBirdsMongoRepository(mongoConn, postgresConn)

	bird, err := r.Get(ctx, "userOwner")
	if err != nil {
		log.Errorf("unable to get bird %v", err)
	}

	require.Equal(t, models.Bird{
		Owner: "userOwner",
		Name:  "Phoenix",
		Type:  "Firebird",
	}, bird)
}

func TestBirdsMongoRepository_Change(t *testing.T) {
	mongoConn, postgresConn := initializeDataBases()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	r := NewBirdsMongoRepository(mongoConn, postgresConn)

	err := r.Change(
		ctx,
		"userOwner",
		models.Bird{Owner: "userOwner", Name: "Dragon", Type: "Oldbird"},
	)
	if err != nil {
		log.Errorf("unable to get bird %v", err)
	}

	require.Nil(t, err)
}

func TestBirdsMongoRepository_Delete(t *testing.T) {
	mongoConn, postgresConn := initializeDataBases()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	r := NewBirdsMongoRepository(mongoConn, postgresConn)

	err := r.Delete(ctx, "userOwner")

	require.Nil(t, err)
}
